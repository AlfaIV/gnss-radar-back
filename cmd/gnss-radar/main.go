package main

import (
	"github.com/Gokert/gnss-radar/configurations"
	"github.com/Gokert/gnss-radar/internal/delivery"
	"github.com/Gokert/gnss-radar/internal/pkg"
	authorization "github.com/Gokert/gnss-radar/internal/service"
	"github.com/Gokert/gnss-radar/internal/store"
	"log"
	"time"
)

func main() {
	postgresConfig, err := configurations.ParsePostgresConfig(utils.utils.PathPostgresConf)
	if err != nil {
		log.Fatalf("configurations.ParsePostgresConfig: %v", err)
		return
	}

	redisConfig, err := configurations.ParseRedisConfig(utils.PathRedisConf)
	if err != nil {
		log.Fatalf("configurations.ParseRedisConfig: %v", err)
		return
	}

	//serviceConfig, err := configurations.ParseServiceConfig(utils.PathServiceConf)
	//if err != nil {
	//	log.Fatalf("configurations.ParseServiceConfig: %v", err)
	//}

	postgresDB, err := store.ConnectToPostgres(postgresConfig)
	if err != nil {
		log.Fatalf("store.ConnectToPostgres: %v", err)
	}
	log.Printf("Successfully connected to postgres")

	redisDB, err := store.ConnectToRedis(redisConfig)
	if err != nil {
		log.Fatalf("store.ConnectToRedis: %v", err)
	}
	log.Printf("Successfully connected to redis")

	storage := store.NewStorage(postgresDB, redisDB)
	storageManager := store.NewStore(storage)

	authorization.NewService(storageManager.GetAuthorizationStore())

	app := delivery.App{}

	log.Printf("App created")
	time.Sleep(10 * time.Second)
}
