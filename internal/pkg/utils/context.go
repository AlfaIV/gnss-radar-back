package utils

import (
	"context"
	"net/http"
)

type contextKey string

const ResponseWriterKey contextKey = "responseWriter"
const RequestKey contextKey = "requestKey"

func SetCookie(ctx context.Context, value string) {
	http.SetCookie(ctx.Value(ResponseWriterKey).(http.ResponseWriter), &http.Cookie{
		Name:     "session_id",
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   60 * 60 * 24,
	})
}

func GetRequest(ctx context.Context) *http.Request {
	request, _ := ctx.Value(RequestKey).(*http.Request)
	return request
}
