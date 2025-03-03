package main

import (
	"context"
	"gnss-radar/gnss-api-gateway/internal/config"
	"gnss-radar/gnss-api-gateway/internal/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	auth_proto "gnss-radar/api/proto/auth"
	auth_handler "gnss-radar/gnss-api-gateway/internal/auth/delivery"
	auth_client "gnss-radar/gnss-api-gateway/internal/auth/service/client"

	user_proto "gnss-radar/api/proto/user"
	user_handler "gnss-radar/gnss-api-gateway/internal/user/delivery"
	user_client "gnss-radar/gnss-api-gateway/internal/user/service/client"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Info("[GATEWAY]: logger initialized")

	// Инициализируем авторизацию

	authGRPCClientConn, err := grpc.NewClient(os.Getenv("AUTH_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("[GATEWAY]: ", err)
	}
	defer authGRPCClientConn.Close()

	authClient := auth_proto.NewAuthClient(authGRPCClientConn)
	authUsecase := auth_client.NewAuthClient(authClient, logger)

	// Инициализируем пользовательский микросервис

	userGRPCClientConn, err := grpc.NewClient(os.Getenv("USER_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("[GATEWAY]: ", err)
	}
	defer userGRPCClientConn.Close()

	userClient := user_proto.NewUserServiceClient(userGRPCClientConn)
	userUsecase := user_client.NewUserClient(userClient, logger)

	// Обработчики на гейтвее

	authHandler := auth_handler.NewHandler(&userUsecase, &authUsecase, logger)
	userHandler := user_handler.NewHandler(&userUsecase, logger)

	config, err := config.NewConfig()
	if err != nil {
		logger.Fatal("[GATEWAY]: ", err)
		os.Exit(1)
	}
	logger.Info("[GATEWAY]: config set up initialized")

	//Прокидывание обработчиков и клиентов для микросервиса в маршрутизатор

	e := mux.Setup(config, mux.ServiceUsecase{
		Auth: &authUsecase,
		User: &userUsecase,
	}, mux.Handlers{
		Auth: authHandler,
		User: userHandler,
	}, logger)

	server := &http.Server{
		Addr:    os.Getenv("GATEWAY_ADDR"),
		Handler: e,
	}

	// Плавная остановка
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("[GATEWAY]: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("[GATEWAY]: Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Errorf("[GATEWAY]: %v", err)
	}

}
