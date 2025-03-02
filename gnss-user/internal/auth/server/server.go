package auth_server

import (
	"context"
	proto "gnss-radar/api/proto/user"

	user_domain "gnss-radar/gnss-user/internal"
	statistics_repository "gnss-radar/gnss-user/internal/repository"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	logger *logrus.Logger
	repo   *statistics_repository.UserRepo

	proto.UnimplementedUserServiceServer
}

func NewAuthServer(repo *statistics_repository.UserRepo, logger *logrus.Logger) UserServiceServer {
	return UserServiceServer{repo: repo, logger: logger}
}

func (s *UserServiceServer) Login(ctx context.Context, req *user.LoginRequest) (*user.User, error) {

	userInfoReq := user_domain.UserInfoRequest{
		Login:    req.Username,
		Password: req.Password,
	}

	userInfo, err := s.repo.GetUserInfo(ctx, userInfoReq)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials: %v", err)
	}

	return &user.User{
		Login:            userInfo.Login,
		Role:             userInfo.Role,
		OrganizationName: userInfo.OrganizationName,
		Name:             userInfo.Name,
		Surname:          userInfo.Surname,
		Email:            userInfo.Email,
		Api:              userInfo.Api,
	}, nil
}

func (s *UserServiceServer) SignUp(ctx context.Context, req *user.SignUpRequest) (*user.Status, error) {

	createUserReq := user_domain.CreateUserRequest{
		Login:            req.Login,
		OrganizationName: req.OrganizationName,
		Name:             req.Name,
		Surname:          req.Surname,
		Password:         req.Password,
		Email:            req.Email,
	}

	if err := s.repo.CreateUser(ctx, createUserReq); err != nil {
		return &user.Status{IsOk: false}, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	return &user.Status{IsOk: true}, nil
}

func (s *UserServiceServer) GetUserInfoById(ctx context.Context, req *user.UserId) (*user.User, error) {

	userInfoReq := user_domain.UserInfoRequest{
		Login: req.Id,
	}

	userInfo, err := s.repo.GetUserInfoById(ctx, userInfoReq) // Дописать для БД
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	return &user.User{
		Login:            userInfo.Login,
		Role:             userInfo.Role,
		OrganizationName: userInfo.OrganizationName,
		Name:             userInfo.Name,
		Surname:          userInfo.Surname,
		Email:            userInfo.Email,
		Api:              userInfo.Api,
	}, nil
}

func (s *UserServiceServer) GetListUsers(ctx context.Context, req *user.PaginatedRequest) (*user.UserList, error) {

	users, err := s.repo.GetUserForAdmin(ctx, user_domain.PaginatedRequest{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user list: %v", err)
	}

	var userList []*user.User
	for _, user := range users {
		userList = append(userList, &user.User{
			Login:            user.Login,
			Role:             user.Role,
			OrganizationName: user.OrganizationName,
			Name:             user.Name,
			Surname:          user.Surname,
			Email:            user.Email,
		})
	}

	return &user.UserList{Users: userList}, nil
}
