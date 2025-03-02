package user_domain_gateway

type UserForAdmin struct {
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Email            string `json:"email"`
	Login            string `json:"login"`
	OrganizationName string `json:"organizationName"`
}

type UserListResponse struct {
	Users []UserForAdmin `json:"users"`
}
