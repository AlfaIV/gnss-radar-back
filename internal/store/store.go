package store

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Gokert/gnss-radar/configurations"
	"github.com/Gokert/gnss-radar/internal/pkg"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"time"
)

type Store struct {
	authorization *AuthorizationStore
}

func NewStore(storage *Storage) *Store {
	return &Store{
		authorization: NewAuthorizationStore(storage),
	}
}

type Storage struct {
	postgres *sql.DB
	redis    *redis.Client
}

func NewStorage(postgres *sql.DB, redis *redis.Client) *Storage {
	return &Storage{
		postgres: postgres,
		redis:    redis,
	}
}

func ConnectToPostgres(config *configurations.DbPsxConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password= %s host=%s port=%d sslmode=%s",
		config.User, config.Dbname, config.Password, config.Host, config.Port, config.Sslmode)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %s", err.Error())
	}
	fmt.Print(config)

	errs := make(chan error)
	go func() {
		errs <- pingDb(db)
	}()

	if err = <-errs; err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenConns)

	return db, nil
}

func ConnectToRedis(config *configurations.DbRedisCfg) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DbNumber,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("ping redis error: %s", err.Error())
	}

	return redisClient, nil
}

func pingDb(db *sql.DB) error {
	var err error
	var retries int

	for retries < utils.MaxRetries {
		err = db.Ping()
		if err == nil {
			return nil
		}

		retries++
		log.Printf("sql ping error: %v", err)

		time.Sleep(time.Duration(utils.MaxTimer) * time.Second)
	}

	return fmt.Errorf("sql max pinging error: %v", err)
}

func (s *Store) GetAuthorizationStore() *AuthorizationStore {
	return s.authorization
}
