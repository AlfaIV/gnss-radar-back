package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Gokert/gnss-radar/configurations"
	"github.com/Gokert/gnss-radar/internal/pkg/consts"
	"github.com/Gokert/gnss-radar/internal/pkg/executor"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Store struct {
	authorization *AuthorizationStore
	session       *SessionStore
	gnss          *GnssStore
}

func NewStore(storage *Storage, cacheStorage *CacheStorage) *Store {
	return &Store{
		authorization: NewAuthorizationStore(storage),
		session:       NewSessionStore(cacheStorage),
		gnss:          NewGnssStore(storage),
	}
}

type Storage struct {
	db executor.Executor
}

type CacheStorage struct {
	db *redis.Client
}

func NewStorage(postgres *pgxpool.Pool) *Storage {
	return &Storage{
		db: executor.NewExecutor(postgres),
	}
}

func NewCacheStorage(redis *redis.Client) *CacheStorage {
	return &CacheStorage{
		db: redis,
	}
}

func (s *Storage) Builder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

type connectResult struct {
	pool *pgxpool.Pool
	err  error
}

func ConnectToPostgres(config *configurations.DbPsxConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password= %s host=%s port=%d sslmode=%s",
		config.User, config.Dbname, config.Password, config.Host, config.Port, config.Sslmode)

	result := make(chan connectResult)
	go func() {
		result <- pingDb(dsn)
	}()

	res := <-result
	if res.err != nil {
		return nil, res.err
	}

	return res.pool, nil
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

func pingDb(dsn string) connectResult {
	var err error
	var retries int

	for retries < consts.MaxRetries {
		connect, err := pgxpool.Connect(context.Background(), dsn)
		if err == nil {
			return connectResult{connect, err}
		}

		retries++
		log.Printf("sql ping error: %v", err)

		time.Sleep(time.Duration(consts.MaxTimer) * time.Second)
	}

	return connectResult{nil, fmt.Errorf("pgxpool.Connect: sql max pinging error: %v", err)}
}

func (s *Store) GetAuthorizationStore() *AuthorizationStore {
	return s.authorization
}

func (s *Store) GetSessionStore() *SessionStore {
	return s.session
}

func (s *Store) GetGnssStore() *GnssStore {
	return s.gnss
}
