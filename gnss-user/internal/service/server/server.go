package auth_server

import (
	"context"
	proto "gnss-radar/api/proto/user"

	user_domain "gnss-radar/gnss-user/internal"

	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	logger *logrus.Logger
	repo   user_domain.Repository

	proto.UnimplementedUserServiceServer
}

func NewUserServer(repo user_domain.Repository, logger *logrus.Logger) UserServiceServer {
	return UserServiceServer{repo: repo, logger: logger}
}

func (s *UserServiceServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.User, error) {

	userInfoReq := user_domain.UserInfoRequest{
		Login:    req.Login,
		Password: req.Password,
	}

	userInfo, err := s.repo.GetUserInfo(ctx, userInfoReq)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "[USER]: %v", err)
	}

	return &proto.User{
		Id:               userInfo.Id,
		Login:            userInfo.Login,
		Role:             userInfo.Role,
		OrganizationName: userInfo.OrganizationName,
		Name:             userInfo.Name,
		Surname:          userInfo.Surname,
		Email:            userInfo.Email,
		Status:           userInfo.Status,
		Api:              userInfo.Api,
	}, nil
}

func (s *UserServiceServer) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.Status, error) {

	createUserReq := user_domain.CreateUserRequest{
		Login:            req.Login,
		OrganizationName: req.OrganizationName,
		Name:             req.Name,
		Surname:          req.Surname,
		Password:         req.Password,
		Email:            req.Email,
	}

	if err := s.repo.CreateUser(ctx, createUserReq); err != nil {
		return &proto.Status{IsOk: false}, status.Errorf(codes.Internal, "[USER]: %v", err)
	}

	return &proto.Status{IsOk: true}, nil
}

func (s *UserServiceServer) GetUserInfoById(ctx context.Context, req *proto.UserId) (*proto.User, error) {

	userInfo, err := s.repo.GetUserInfoById(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "[USER]: %v", err)
	}

	return &proto.User{
		Id:               userInfo.Id,
		Login:            userInfo.Login,
		Role:             userInfo.Role,
		OrganizationName: userInfo.OrganizationName,
		Name:             userInfo.Name,
		Surname:          userInfo.Surname,
		Email:            userInfo.Email,
		Status:           userInfo.Status,
		Api:              userInfo.Api,
	}, nil
}

func (s *UserServiceServer) GetListUsers(ctx context.Context, req *proto.PaginatedRequest) (*proto.UserList, error) {

	users, err := s.repo.GetUserForAdmin(ctx, user_domain.PaginatedRequest{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "[USER]: %v", err)
	}

	var userList []*proto.User
	for _, user := range users {
		userList = append(userList, &proto.User{
			Login:            user.Login,
			Role:             user.Role,
			OrganizationName: user.OrganizationName,
			Name:             user.Name,
			Surname:          user.Surname,
			Email:            user.Email,
		})
	}

	return &proto.UserList{Users: userList}, nil
}

func (s *UserServiceServer) GetSignUpRequestions(ctx context.Context, req *proto.PaginatedRequest) (*proto.UserList, error) {

	users, err := s.repo.GetSignUpRequestions(ctx, user_domain.PaginatedRequest{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "[USER]: %v", err)
	}

	var userList []*proto.User
	for _, user := range users {
		userList = append(userList, &proto.User{
			Login:            user.Login,
			OrganizationName: user.OrganizationName,
			Name:             user.Name,
			Surname:          user.Surname,
			Email:            user.Email,
		})
	}

	return &proto.UserList{Users: userList}, nil
}

func (s *UserServiceServer) ValidatePermissions(ctx context.Context, req *proto.PermissionValidaton) (*proto.Status, error) {

	result, err := s.repo.ValidatePermissions(ctx, req.UserId, req.Api)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "[USER]: %v", err)
	}

	return &proto.Status{IsOk: result}, nil
}

func (s *UserServiceServer) ResolveUserSignUp(ctx context.Context, req *proto.SignUpResolution) (*google_proto.Empty, error) {
	err := s.repo.ResolveUserSignUp(ctx, req.UserLogin, req.Resolution)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "[USER]: %v", err)
	}

	return &google_proto.Empty{}, nil
}

func (s *UserServiceServer) ChangeUserPermissions(ctx context.Context, req *proto.PermissionChange) (*google_proto.Empty, error) {
	err := s.repo.ResolveUserSignUp(ctx, req.UserLogin, req.UserRole)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "[USER]: %v", err)
	}

	return &google_proto.Empty{}, nil
}
