package auth_client

import (
	"context"
	proto "gnss-radar/api/proto/auth"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type AuthClient struct {
	client proto.AuthClient
	logger *logrus.Logger
}

func NewAuthClient(client proto.AuthClient, logger *logrus.Logger) AuthClient {
	return AuthClient{client: client, logger: logger}
}

func (ac *AuthClient) CheckSession(ctx context.Context, sessionId string) (bool, error) {
	isOk, err := ac.client.CheckSession(ctx, &proto.SessionId{SessionId: sessionId})
	if err != nil {
		return false, errors.Wrapf(err, "failed to check sessionId %s", sessionId)
	}

	return isOk.GetIsOk(), nil
}

func (ac *AuthClient) CreateSession(ctx context.Context, userId string) (string, error) {
	sessionId, err := ac.client.CreateSession(ctx, &proto.UserId{UserId: userId})
	if err != nil {
		return "", errors.Wrapf(err, "failed to create session for user %s", userId)
	}

	return sessionId.GetSessionId(), nil
}

func (ac *AuthClient) DeleteSession(ctx context.Context, sessionId string) error {
	if _, err := ac.client.DeleteSession(ctx, &proto.SessionId{SessionId: sessionId}); err != nil {
		return errors.Wrapf(err, "failed to delete session %s", sessionId)
	}

	return nil
}

func (ac *AuthClient) GetUserId(ctx context.Context, sessionId string) (string, error) {
	userId, err := ac.client.GetUserId(ctx, &proto.SessionId{SessionId: sessionId})
	if err != nil {
		return "", errors.Wrapf(err, "failed to get user by sessionId %s", sessionId)
	}

	return userId.GetUserId(), nil
}
