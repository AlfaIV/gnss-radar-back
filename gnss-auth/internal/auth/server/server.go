package auth_server

import (
	"context"
	proto "gnss-radar/api/proto/auth"
	common_proto "gnss-radar/api/proto/common"
	auth_domain "gnss-radar/gnss-auth/internal/auth"

	google_proto "github.com/golang/protobuf/ptypes/empty"

	"github.com/sirupsen/logrus"
)

type AuthServer struct {
	repo   auth_domain.Repository
	logger *logrus.Logger
	proto.UnimplementedAuthServer
}

func NewAuthServer(repo auth_domain.Repository, logger *logrus.Logger) AuthServer {
	return AuthServer{repo: repo, logger: logger}
}

func (as *AuthServer) CheckSession(ctx context.Context, in *proto.SessionId) (*common_proto.Status, error) {
	if _, err := as.repo.GetId(ctx, in.GetSessionId()); err != nil {
		return &common_proto.Status{IsOk: false}, err
	}

	return &common_proto.Status{IsOk: true}, nil
}

func (as *AuthServer) CreateSession(ctx context.Context, in *common_proto.UserId) (*proto.SessionId, error) {
	sessionId, err := as.repo.Set(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	return &proto.SessionId{SessionId: sessionId}, nil
}

func (as *AuthServer) DeleteSession(ctx context.Context, in *proto.SessionId) (*google_proto.Empty, error) {
	if err := as.repo.Delete(ctx, in.GetSessionId()); err != nil {
		return nil, err
	}

	return &google_proto.Empty{}, nil
}

func (as *AuthServer) GetUserId(ctx context.Context, in *proto.SessionId) (*common_proto.UserId, error) {
	userId, err := as.repo.GetId(ctx, in.GetSessionId())
	if err != nil {
		return nil, err
	}

	return &common_proto.UserId{UserId: userId}, nil
}
