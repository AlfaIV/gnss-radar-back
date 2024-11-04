package main

import (
	"log"

	"github.com/Gokert/gnss-radar/configurations"
	"github.com/Gokert/gnss-radar/internal/delivery"
	"github.com/Gokert/gnss-radar/internal/pkg/consts"
	"github.com/Gokert/gnss-radar/internal/service"
	"github.com/Gokert/gnss-radar/internal/store"
)

func main() {
	postgresConfig, err := configurations.ParsePostgresConfig(consts.PathPostgresConf)
	if err != nil {
		log.Fatalf("configurations.ParsePostgresConfig: %v", err)
		return
	}

	redisConfig, err := configurations.ParseRedisConfig(consts.PathRedisConf)
	if err != nil {
		log.Fatalf("configurations.ParseRedisConfig: %v", err)
		return
	}

	serviceConfig, err := configurations.ParseServiceConfig(consts.PathServiceConf)
	if err != nil {
		log.Fatalf("configurations.ParseServiceConfig: %v", err)
	}

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

	storage := store.NewStorage(postgresDB)
	cacheStorage := store.NewCacheStorage(redisDB)

	storageManager := store.NewStore(storage, cacheStorage)

	newService := service.NewService(storageManager.GetAuthorizationStore(), storageManager.GetSessionStore(), storageManager.GetGnssStore())
	app := delivery.NewApp(newService)

	//
	// Run app
	//
	if err = app.Run(serviceConfig.Port); err != nil {
		log.Fatalf("delivery.Run: %v", err)
	}
}
