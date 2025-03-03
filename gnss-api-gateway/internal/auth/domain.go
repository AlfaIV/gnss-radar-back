package auth_domain_gateway

import (
	"context"
	"time"
)

type Usecase interface {
	CheckSession(ctx context.Context, sessionId string) (bool, error)
	CreateSession(ctx context.Context, userId string) (string, error)
	DeleteSession(ctx context.Context, sessionId string) error
	GetUserId(ctx context.Context, sessionId string) (string, error)
}

const CookieName = "SessionID"
const CookieTTL = 24 * 14 * time.Hour