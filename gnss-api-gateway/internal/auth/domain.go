package auth_domain

import (
	"context"
)

type Usecase interface {
	CheckSession(ctx context.Context, sessionId string) (bool, error)
	CreateSession(ctx context.Context, userId string) (string, error)
	DeleteSession(ctx context.Context, sessionId string) error
	GetUserId(ctx context.Context, sessionId string) (string, error)
}

const CookieName = "SessionID"
