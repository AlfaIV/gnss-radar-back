package main

import (
	"context"
	auth_proto "gnss-radar/api/proto/auth"
	auth_server "gnss-radar/gnss-auth/internal/auth/server"
	auth_repository "gnss-radar/gnss-auth/internal/repository"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Info("[AUTH]: logger initialized")

	lis, err := net.Listen("tcp", os.Getenv("AUTH_ADDR"))
	if err != nil {
		logger.Fatal("[AUTH]: ", err)
		os.Exit(1)
	}

	// Инициализация Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	// Проверка подключения к Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		logger.Fatal("[AUTH]: Redis connection failed - ", err)
	}
	logger.Info("[AUTH]: Redis initialized")

	// Инициализация сервисов
	authRepository := auth_repository.NewAuth(rdb, logger)
	authServer := auth_server.NewAuthServer(authRepository, logger)

	// Создание gRPC сервера
	server := grpc.NewServer()
	
	// Регистрация health check сервиса
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, healthServer)
	
	// Установка начального статуса
	healthServer.SetServingStatus("auth.Service", grpc_health_v1.HealthCheckResponse_SERVING)

	// Регистрация основного сервиса
	auth_proto.RegisterAuthServer(server, &authServer)

	// Graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh

		logger.Info("[AUTH]: Starting graceful shutdown...")
		
		// Пометить сервис как NOT_SERVING
		healthServer.SetServingStatus("auth.Service", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
		
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
		
		// Закрытие соединений
		rdb.Close()
		lis.Close()
		
		logger.Info("Server stopped gracefully")
		os.Exit(0)
	}()

	logger.Info("[AUTH]: starting auth server at ", os.Getenv("AUTH_ADDR"))
	if err := server.Serve(lis); err != nil {
		logger.Fatal("[AUTH]: failed to serve - ", err)
	}
}