package middleware

import (
	"context"
	"github.com/Gokert/gnss-radar/internal/pkg/utils"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

var Middlewares = []Middleware{
	SetResponseRequest,
}

func CallMiddlewares(middlewares ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		for _, m := range middlewares {
			final = m(final)
		}
		return final
	}
}

func SetResponseRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), utils.ResponseWriterKey, w)
		ctx = context.WithValue(ctx, utils.RequestKey, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
