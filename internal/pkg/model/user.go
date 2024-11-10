package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Login     string    `json:"login"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
