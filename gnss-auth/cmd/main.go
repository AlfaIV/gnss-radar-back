package main

import (
	"context"
	auth_proto "gnss-radar/api/proto/auth"
	auth_server "gnss-radar/gnss-auth/internal/auth/server"
	auth_repository "gnss-radar/gnss-auth/internal/repository"
	"net"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		logger.Fatal("[AUTH]: ", err)
		os.Exit(1)
	}

	logger.Info("[AUTH]: redis initialized")

	authRepository := auth_repository.NewAuth(rdb, logger)
	authServer := auth_server.NewAuthServer(authRepository, logger)

	server := grpc.NewServer()
	auth_proto.RegisterAuthServer(server, &authServer)

	logger.Info("[AUTH]: starting auth server at ", os.Getenv("AUTH_ADDR"))
	logger.Fatal(server.Serve(lis))
}
