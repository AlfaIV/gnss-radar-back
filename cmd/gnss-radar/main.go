package main

import (
	"log"
	"sync"

	middleware2 "github.com/Gokert/gnss-radar/internal/pkg/middleware"

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
	authService := service.NewService(storageManager.GetAuthorizationStore(), storageManager.GetSessionStore(), storageManager.GetGnssStore())
	middleware := middleware2.NewService(authService.GetAuthorizationService())

	hardwareService := service.NewHardwareService(storageManager.GetGnssStore())
	graphqlApp := delivery.NewApp(authService, hardwareService, middleware)

	grpcListenerServer, err := delivery.NewServer()
	if err != nil {
		log.Fatalf("delivery.NewServer: %v", err)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(3)

	//
	// Graphql app run
	//
	go func() {
		defer wg.Done()
		if err = graphqlApp.Run(serviceConfig.GraphqlPort); err != nil {
			log.Fatalf("graphqlApp.Run: %v", err)
			return
		}
	}()

	//
	// Grpc listener run
	//
	go func() {
		defer wg.Done()
		if err = grpcListenerServer.ListenAndServeGrpc(serviceConfig.ConnectionType, serviceConfig.GrpcListenerPort); err != nil {
			log.Fatalf("grpcListenerServer.ListenAndServeGrpc: %v", err)
			return
		}
	}()

	//
	// Hardware rest app run
	//
	go func() {
		defer wg.Done()
		if err := graphqlApp.HardwareHandlers(serviceConfig.GraphqlPort); err != nil {
			log.Fatalf("graphqlApp.HardwareHandlers: %v", err)
			return
		}
	}()

	wg.Wait()
}
