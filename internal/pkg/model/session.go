package model

import "time"

type Session struct {
	Login     string
	SID       string
	ExpiresAt time.Time
}
