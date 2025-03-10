package mwutils

import (
	"context"
	middlewarecustom "gnss-radar/gnss-api-gateway/internal/mux/middleware"
	"net/http"
)

type AppError struct {
	CodeHTTP int
	Message  string
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	UnauthorizedInHandler = AppError{CodeHTTP: http.StatusInternalServerError, Message: "unauthorized request made it past auth check middleware"}
	IDConversion          = AppError{CodeHTTP: http.StatusInternalServerError, Message: "invalid ID fetched for current session (must be uuid)"}
)

func GetUserID(ctx context.Context) (string, error) {
	contextUserId := ctx.Value(middlewarecustom.UserIDKey)
	if contextUserId == nil {
		return "", &UnauthorizedInHandler
	}

	userIdStr, ok := contextUserId.(string)
	if !ok {
		return "", &IDConversion
	}

	return userIdStr, nil
}
