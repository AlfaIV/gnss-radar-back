package model

type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Role  string `json:"role"`
}
