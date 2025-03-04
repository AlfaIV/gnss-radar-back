package auth_repository

import (
	"context"
	"fmt"
	auth_domain "gnss-radar/gnss-auth/internal/auth"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	db     *redis.Client
	logger *logrus.Logger
}

func NewAuth(db *redis.Client, logger *logrus.Logger) *Auth {
	return &Auth{db: db, logger: logger}
}

func (redis *Auth) Set(ctx context.Context, userId string) (string, error) {
	sessionId := uuid.New().String()
	fmt.Println(userId)

	if err := redis.db.Set(ctx, sessionId, userId, auth_domain.CookieTTL).Err(); err != nil {
		redis.logger.Error("[AUTH]:", err.Error())
		return "", err
	}
	redis.logger.WithField("sessionId", sessionId).Info("[AUTH]: session created")

	return sessionId, nil
}

func (redis *Auth) Delete(ctx context.Context, sessionId string) error {
	if err := redis.db.Del(ctx, sessionId).Err(); err != nil {
		redis.logger.Error("[AUTH]:", err.Error())
		return err
	}

	redis.logger.WithField("sessionId", sessionId).Info("[AUTH]: session deleted")

	return nil
}

func (redis *Auth) GetId(ctx context.Context, sessionId string) (string, error) {
	userId, err := redis.db.Get(ctx, sessionId).Result()
	if err != nil {
		redis.logger.Error("[AUTH]:", err.Error())
		return "", err
	}

	redis.logger.WithField("userId", userId).Info("[AUTH]: got user id")

	return userId, nil
}
