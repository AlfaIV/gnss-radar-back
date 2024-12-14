package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/go-redis/redis/v8"
	"time"
)

type ISessionStore interface {
	AddSession(ctx context.Context, active model.Session) (bool, error)
	CheckActiveSession(ctx context.Context, sid string) (bool, error)
	GetUserLogin(ctx context.Context, sid string) (string, error)
	DeleteSession(ctx context.Context, sid string) (bool, error)
}

type SessionStore struct {
	sessionStorage *CacheStorage
}

func NewSessionStore(cache *CacheStorage) *SessionStore {
	return &SessionStore{
		sessionStorage: cache,
	}
}

func (s *SessionStore) AddSession(ctx context.Context, session model.Session) (bool, error) {
	if _, err := s.sessionStorage.db.Set(ctx, session.SID, session.Login, 24*time.Hour).Result(); err != nil {
		return false, fmt.Errorf(".sessionStorage.db.Set: %w", err)
	}

	return true, nil
}

func (s *SessionStore) CheckActiveSession(ctx context.Context, sid string) (bool, error) {
	_, err := s.sessionStorage.db.Get(ctx, sid).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, fmt.Errorf("sessionStorage.db.Get: %w", err)
	}

	return true, nil
}

func (s *SessionStore) GetUserLogin(ctx context.Context, sid string) (string, error) {
	value, err := s.sessionStorage.db.Get(ctx, sid).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrNotFound
		}
		return "", fmt.Errorf("sessionStorage.db.Get: %w", err)
	}

	return value, nil
}

func (s *SessionStore) DeleteSession(ctx context.Context, sid string) (bool, error) {
	if _, err := s.sessionStorage.db.Del(ctx, sid).Result(); err != nil {
		return false, fmt.Errorf("s.sessionStorage.db.Del: %w", err)
	}

	return true, nil
}
