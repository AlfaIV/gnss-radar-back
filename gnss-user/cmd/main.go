package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	user_proto "gnss-radar/api/proto/user"
	user_repository "gnss-radar/gnss-user/internal/repository"
	user_server "gnss-radar/gnss-user/internal/service/server"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Info("[USER]: logger initialized")

	// Создание TCP-листенера
	lis, err := net.Listen("tcp", os.Getenv("USER_ADDR"))
	if err != nil {
		logger.Fatal("[USER]: failed to listen: ", err)
		os.Exit(1)
	}

	ctx := context.Background()
	// Парсинг конфигурации PostgreSQL
	configpg, err := pgxpool.ParseConfig(os.Getenv("PG_ADDR"))
	if err != nil {
		logger.Fatal("[USER]: failed to parse postgres config: ", err)
		os.Exit(1)
	}

	configpg.MaxConns = 20
	// Инициализация пула подключений к PostgreSQL
	connPool, err := pgxpool.NewWithConfig(ctx, configpg)
	if err != nil {
		logger.Fatal("[USER]: failed to create postgres connection pool: ", err)
		os.Exit(1)
	}
	logger.Info("[USER]: postgres initialized")

	// Инициализация репозитория и сервера пользователя
	userRepository := user_repository.NewUserRepo(connPool, logger)
	userServer := user_server.NewUserServer(userRepository, logger)

	// Создание gRPC сервера
	server := grpc.NewServer()

	// Регистрация health check сервиса
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, healthServer)
	healthServer.SetServingStatus("user.Service", grpc_health_v1.HealthCheckResponse_SERVING)

	// Регистрация основного сервиса
	user_proto.RegisterUserServiceServer(server, &userServer)

	// Graceful shutdown: обработка сигналов завершения
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh

		logger.Info("[USER]: starting graceful shutdown...")

		// Пометить сервис как NOT_SERVING
		healthServer.SetServingStatus("user.Service", grpc_health_v1.HealthCheckResponse_NOT_SERVING)

		// Остановка сервера с таймаутом
		stopped := make(chan struct{})
		go func() {
			server.GracefulStop()
			close(stopped)
		}()

		select {
		case <-time.After(15 * time.Second):
			server.Stop()
		case <-stopped:
		}

		// Закрытие соединения с PostgreSQL и TCP-листенера
		connPool.Close()
		lis.Close()

		logger.Info("[USER]: server stopped gracefully")
		os.Exit(0)
	}()

	logger.Info("[USER]: starting user server at ", os.Getenv("USER_ADDR"))
	if err := server.Serve(lis); err != nil {
		logger.Fatal("[USER]: failed to serve: ", err)
	}
}
