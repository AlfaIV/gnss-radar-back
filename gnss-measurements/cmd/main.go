package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	measurements_proto "gnss-radar/api/proto/measurements"
	measurements_repository "gnss-radar/gnss-measurements/internal/repository"
	measurements_server "gnss-radar/gnss-measurements/internal/service/server"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Info("[MEASUREMENTS]: logger initialized")

	// Создание TCP-листенера
	lis, err := net.Listen("tcp", os.Getenv("MEASUREMENTS_ADDR"))
	if err != nil {
		logger.Fatal("[MEASUREMENTS]: failed to listen: ", err)
		os.Exit(1)
	}

	minioClient, err := minio.New(os.Getenv("MINIO_API_ADDR"), &minio.Options{
		Creds: credentials.NewStaticV4(
			os.Getenv("MINIO_ROOT_USER"),
			os.Getenv("MINIO_ROOT_PASSWORD"),
			"",
		),
		Secure: false,
	})
	if err != nil {
		logger.Fatal("[MEASUREMENTS]: ", err)
		os.Exit(1)
	}

	ctx := context.Background()

	//Проверка нужных бакетов для S3
	requiredBuckets := []string{os.Getenv("EPHEMERIS_BUCKET_NAME")}

	for _, bucketName := range requiredBuckets {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists != nil {
			logger.Fatal("[MEASUREMENTS]: failed to check if bucket exists: ", errBucketExists)
			os.Exit(1)
		}
		if !exists {
			logger.Fatalf("[MEASUREMENTS]: bucket %s does not exist", bucketName)
			os.Exit(1)
		}
	}
	logger.Info("[MEASUREMENTS]: buckets checked")

	// Парсинг конфигурации PostgreSQL
	configpg, err := pgxpool.ParseConfig(os.Getenv("PG_ADDR"))
	if err != nil {
		logger.Fatal("[MEASUREMENTS]: failed to parse postgres config: ", err)
		os.Exit(1)
	}

	configpg.MaxConns = 20
	// Инициализация пула подключений к PostgreSQL
	connPool, err := pgxpool.NewWithConfig(ctx, configpg)
	if err != nil {
		logger.Fatal("[MEASUREMENTS]: failed to create postgres connection pool: ", err)
		os.Exit(1)
	}
	logger.Info("[MEASUREMENTS]: postgres initialized")

	// Инициализация репозитория и сервера пользователя
	measurementsRepo := measurements_repository.NewMeasurementsRepo(connPool, minioClient, logger)
	measurementsServer := measurements_server.NewMeasurementsServer(measurementsRepo, logger)

	// Создание gRPC сервера
	server := grpc.NewServer()

	// Регистрация health check сервиса
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, healthServer)
	healthServer.SetServingStatus("measurements.Service", grpc_health_v1.HealthCheckResponse_SERVING)

	// Регистрация основного сервиса
	measurements_proto.RegisterMeasurementsServer(server, &measurementsServer)

	// Graceful shutdown: обработка сигналов завершения
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh

		logger.Info("[MEASUREMENTS]: starting graceful shutdown...")

		// Пометить сервис как NOT_SERVING
		healthServer.SetServingStatus("measurements.Service", grpc_health_v1.HealthCheckResponse_NOT_SERVING)

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

		logger.Info("[MEASUREMENTS]: server stopped gracefully")
		os.Exit(0)
	}()

	logger.Info("[MEASUREMENTS]: starting user server at ", os.Getenv("MEASUREMENTS_ADDR"))
	if err := server.Serve(lis); err != nil {
		logger.Fatal("[MEASUREMENTS]: failed to serve: ", err)
	}
}
