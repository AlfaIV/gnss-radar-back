package user_client

import (
	"context"
	proto "gnss-radar/api/proto/user"
	user_domain_gateway "gnss-radar/gnss-api-gateway/internal/user"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type UserClient struct {
	client proto.UserServiceClient
	logger *logrus.Logger
}

func NewUserClient(client proto.UserServiceClient, logger *logrus.Logger) UserClient {
	return UserClient{client: client, logger: logger}
}

func (uc *UserClient) Login(ctx context.Context, login string, password string) (user_domain_gateway.User, error) {
	user, err := uc.client.Login(ctx, &proto.LoginRequest{Login: login, Password: password})
	if err != nil {
		return user_domain_gateway.User{}, errors.Wrapf(err, "[GW USER] %v", err)
	}

	return user_domain_gateway.User{
		Id:               user.Id,
		Name:             user.Name,
		Surname:          user.Surname,
		Email:            user.Email,
		Login:            user.Login,
		OrganizationName: user.OrganizationName,
		Role:             user.Role,
		Status:           user.Status,
		Api:              user.Api,
	}, nil
}

func (uc *UserClient) SignUp(ctx context.Context, req user_domain_gateway.SignUpRequest) (bool, error) {
	status, err := uc.client.SignUp(ctx, &proto.SignUpRequest{
		Login:            req.Login,
		OrganizationName: req.OrganizationName,
		Name:             req.Name,
		Surname:          req.Surname,
		Password:         req.Password,
		Email:            req.Email,
		Role:             req.Role,
	})
	if err != nil {
		return false, errors.Wrapf(err, "[GW USER] %v", err)
	}

	return status.IsOk, nil
}

func (uc *UserClient) GetUserInfoById(ctx context.Context, userId string) (user_domain_gateway.User, error) {
	user, err := uc.client.GetUserInfoById(ctx, &proto.UserId{Id: userId})
	if err != nil {
		return user_domain_gateway.User{}, errors.Wrapf(err, "[GW USER] %v", err)
	}

	return user_domain_gateway.User{
		Id:               user.Id,
		Name:             user.Name,
		Surname:          user.Surname,
		Email:            user.Email,
		Login:            user.Login,
		OrganizationName: user.OrganizationName,
		Role:             user.Role,
		Status:           user.Status,
		Api:              user.Api,
	}, nil
}

func (uc *UserClient) GetListUsers(ctx context.Context, page uint64, size uint64) (user_domain_gateway.UserListResponse, error) {
	users, err := uc.client.GetListUsers(ctx, &proto.PaginatedRequest{Page: page, Size: size})
	if err != nil {
		return user_domain_gateway.UserListResponse{}, errors.Wrapf(err, "[GW USER] %v", err)
	}

	var userArray []user_domain_gateway.UserForAdmin

	for _, u := range users.Users {
		userArray = append(userArray, user_domain_gateway.UserForAdmin{
			Name:             u.Name,
			Surname:          u.Surname,
			Email:            u.Email,
			Login:            u.Login,
			OrganizationName: u.OrganizationName,
			Role:             u.Role,
		})
	}

	return user_domain_gateway.UserListResponse{Users: userArray}, nil
}

func (uc *UserClient) GetSignUpRequestions(ctx context.Context, page uint64, size uint64) (user_domain_gateway.UserListResponse, error) {
	users, err := uc.client.GetListUsers(ctx, &proto.PaginatedRequest{Page: page, Size: size})
	if err != nil {
		return user_domain_gateway.UserListResponse{}, errors.Wrapf(err, "[GW USER] %v", err)
	}

	var userArray []user_domain_gateway.UserForAdmin

	for _, u := range users.Users {
		userArray = append(userArray, user_domain_gateway.UserForAdmin{
			Name:             u.Name,
			Surname:          u.Surname,
			Email:            u.Email,
			Login:            u.Login,
			OrganizationName: u.OrganizationName,
			Role:             u.Role,
		})
	}

	return user_domain_gateway.UserListResponse{Users: userArray}, nil
}

//Только для миддлваря
func (uc *UserClient) ValidatePermissions(ctx context.Context, userId string, api string) (bool, error) {
	status, err := uc.client.ValidatePermissions(ctx, &proto.PermissionValidaton{UserId: userId, Api: api})
	if err != nil {
		return false, errors.Wrapf(err, "[GW USER] %v", err)
	}

	return status.IsOk, nil
}

func (uc *UserClient) ResolveUserSignUp(ctx context.Context, userLogin string, resolution string) error {
	_, err := uc.client.ResolveUserSignUp(ctx, &proto.SignUpResolution{UserLogin: userLogin, Resolution: resolution})
	if err != nil {
		return errors.Wrapf(err, "[GW USER] %v", err)
	}

	return nil
}

func (uc *UserClient) ChangeUserPermissions(ctx context.Context, userLogin string, role string) error {
	_, err := uc.client.ChangeUserPermissions(ctx, &proto.PermissionChange{UserLogin: userLogin, UserRole: role})
	if err != nil {
		return errors.Wrapf(err, "[GW USER] %v", err)
	}

	return nil
}
