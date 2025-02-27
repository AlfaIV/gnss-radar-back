package user_domain

import "context"

type UserInfoRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserInfoResponse struct {
	Id               string   `json:"id"`
	Login            string   `json:"login"`
	Role             string   `json:"role"`
	OrganizationName string   `json:"organizationName"`
	Name             string   `json:"name"`
	Surname          string   `json:"surname"`
	Email            string   `json:"email"`
	Api              []string `json:"api"`
}

type CreateUserRequest struct {
	Login            string `json:"login"`
	OrganizationName string `json:"organizationName"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
}

type Repository interface {
	GetUserInfo(ctx context.Context, request UserInfoRequest) (UserInfoResponse, error)
	CreateUser(ctx context.Context, request CreateUserRequest) error
	ValidatePermissions(ctx context.Context, userId string, api string) (bool, error)
	ResolveUserSignUp(ctx context.Context, userLogin string, resolution string) error
	ChangeUserPermissions(ctx context.Context, userLogin string, userRole string) error
}
