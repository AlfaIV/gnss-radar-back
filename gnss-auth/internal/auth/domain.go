package auth_domain

import (
	"context"
	"time"
)

type Repository interface {
	Set(ctx context.Context, userId string) (string, error)
	Delete(ctx context.Context, sessionId string) error
	GetId(ctx context.Context, sessionId string) (string, error)
}


//2 недели жива сессия
const CookieTTL = 24 * 14 * time.Hour
