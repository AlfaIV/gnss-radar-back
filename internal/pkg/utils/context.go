package utils

import (
	"context"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/samber/lo"
	"log"
	"net/http"
	"time"
)

type contextKey string
type roleKey model.Roles

const ResponseWriterKey contextKey = "responseWriter"
const RequestKey contextKey = "requestKey"
const UserRoleKey roleKey = "userRoleKey"

func SetCookie(ctx context.Context, value string) {
	http.SetCookie(ctx.Value(ResponseWriterKey).(http.ResponseWriter), &http.Cookie{
		Name:     "session_id",
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   60 * 60 * 24,
	})
}

func RemoveCookie(ctx context.Context) {
	req := GetRequest(ctx)
	cookie, err := req.Cookie("session_id")
	if err != nil {
		log.Printf("cookie %v not found", req)
		return
	}

	cookie.Expires = time.Now().AddDate(0, 0, -1)

	w := ctx.Value(ResponseWriterKey).(http.ResponseWriter)
	http.SetCookie(w, cookie)
}

func GetRequest(ctx context.Context) *http.Request {
	request, _ := ctx.Value(RequestKey).(*http.Request)
	return request
}

func GetCookie(ctx context.Context) *string {
	cookie, _ := ctx.Value(ResponseWriterKey).(*http.Cookie)
	if cookie == nil {
		return nil
	}
	return &cookie.Value
}

func CheckPermission(ctx context.Context, needRoles []model.Roles) bool {
	name, _ := ctx.Value(UserRoleKey).(model.Roles)
	if !name.IsValid() {
		return false
	}

	return lo.Contains(needRoles, name)
}
