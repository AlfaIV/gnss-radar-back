package main

import (
	"context"
	"net"
	"os"

	user_proto "gnss-radar/api/proto/user"
	user_repository "gnss-radar/gnss-user/internal/repository"
	user_server "gnss-radar/gnss-user/internal/service/server"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Info("[USER]: logger initialized")

	lis, err := net.Listen("tcp", os.Getenv("USER_ADDR"))
	if err != nil {
		logger.Fatal("[USER]: ", err)
		os.Exit(1)
	}

	ctx := context.Background()

	configpg, err := pgxpool.ParseConfig(os.Getenv("PG_ADDR"))
	if err != nil {
		logger.Fatal("[USER]: ", err)
		os.Exit(1)
	}

	configpg.MaxConns = 20

	connPool, err := pgxpool.NewWithConfig(ctx, configpg)
	if err != nil {
		logger.Fatal("[USER]: ", err)
		os.Exit(1)
	}
	defer connPool.Close()
	logger.Info("[USER]: postgres initialized")

	statRepository := user_repository.NewUserRepo(connPool, logger)
	statServer := user_server.NewUserServer(statRepository, logger)

	server := grpc.NewServer()
	user_proto.RegisterUserServiceServer(server, &statServer)

	logger.Info("[USER]: starting user server at ", os.Getenv("USER_ADDR"))
	logger.Fatal(server.Serve(lis))
}
