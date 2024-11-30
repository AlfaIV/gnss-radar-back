package model

import "time"

type User struct {
	ID               string    `json:"id"`
	Login            string    `json:"login"`
	Role             string    `json:"role"`
	Email            string    `json:"email"`
	OrganizationName string    `json:"organization_name"`
	FirstName        string    `json:"first_name"`
	SecondName       string    `json:"second_name"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
