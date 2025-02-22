package main

import (
	"gnss-radar/gnss-api-gateway/internal/config"
	"gnss-radar/gnss-api-gateway/internal/mux"
	"os"

	auth_proto "gnss-radar/gnss-api-gateway/internal/auth/proto"
	auth_client "gnss-radar/gnss-api-gateway/internal/auth/service/client"

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

	config, err := config.NewConfig()
	if err != nil {
		logger.Fatal("[GATEWAY]: ", err)
		os.Exit(1)
	}
	logger.Info("[GATEWAY]: config set up initialized")

	e := mux.Setup(config, &authUsecase, logger)

	e.Logger.Fatal(e.Start(os.Getenv("GATEWAY_ADDR")))

}