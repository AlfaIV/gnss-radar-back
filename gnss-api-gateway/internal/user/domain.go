package user_domain_gateway

import "context"

type User struct {
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	Surname          string   `json:"surname"`
	Email            string   `json:"email"`
	Login            string   `json:"login"`
	OrganizationName string   `json:"organizationName"`
	Role             string   `json:"role"`
	Status           string   `json:"status"`
	Api              []string `json:"api"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Login            string `json:"login"`
	OrganizationName string `json:"organizationName"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	Role             string `json:"role"`
}

type UserForAdmin struct {
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Email            string `json:"email"`
	Login            string `json:"login"`
	OrganizationName string `json:"organizationName"`
	Role             string `json:"role"`
}

type UserListResponse struct {
	Users []UserForAdmin `json:"users"`
}

type SignUpResolutionRequest struct {
	UserLogin  string `json:"userLogin"`
	Resolution string `json:"resolution"`
}

type PermissionChangeRequest struct {
	UserLogin string `json:"userLogin"`
	NewRole   string `json:"newRole"`
}

type Usecase interface {
	Login(ctx context.Context, login string, password string) (User, error)
	SignUp(ctx context.Context, req SignUpRequest) (bool, error)
	GetUserInfoById(ctx context.Context, userId string) (User, error)
	GetListUsers(ctx context.Context, page uint64, size uint64) (UserListResponse, error)
	GetSignUpRequestions(ctx context.Context, page uint64, size uint64) (UserListResponse, error)
	ValidatePermissions(ctx context.Context, userId string, api string) (bool, error)
	ResolveUserSignUp(ctx context.Context, userLogin string, resolution string) error
	ChangeUserPermissions(ctx context.Context, userLogin string, role string) error
}
