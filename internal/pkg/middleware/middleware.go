package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/pkg/utils"
	"github.com/Gokert/gnss-radar/internal/service"
)

type Middleware func(http.Handler) http.Handler

const permissionDenied = "permission denied"

type IMiddlewareService interface {
	CallMiddlewares() Middleware
	SetResponseRequest(next http.Handler) http.Handler
	SetRole(next http.Handler) http.Handler
	CheckAuthorize(next http.Handler) http.Handler
}

type Service struct {
	authService service.IAuthorizationService
}

func NewService(authService service.IAuthorizationService) IMiddlewareService {
	return &Service{
		authService: authService,
	}
}

func (s *Service) CallMiddlewares() Middleware {
	return func(final http.Handler) http.Handler {
		for _, m := range s.getMiddlewares() {
			final = m(final)
		}
		return final
	}
}

func (s *Service) getMiddlewares() []Middleware {
	return []Middleware{
		s.SetResponseRequest,
		s.SetRole,
	}
}

func (s *Service) SetResponseRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), utils.ResponseWriterKey, w)
		ctx = context.WithValue(ctx, utils.RequestKey, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Service) SetRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if errors.Is(err, http.ErrNoCookie) {
			r = r.WithContext(context.WithValue(r.Context(), utils.UserRoleKey, model.RolesUnknown))
			next.ServeHTTP(w, r)
			return
		}

		_, user, err := s.authService.Authcheck(r.Context(), session.Value)
		if err != nil && user == nil {
			r = r.WithContext(context.WithValue(r.Context(), utils.UserRoleKey, model.RolesUnknown))
			next.ServeHTTP(w, r)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), utils.UserRoleKey, model.Roles(user.Role)))
		next.ServeHTTP(w, r)
	})
}

func (s *Service) CheckAuthorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if errors.Is(err, http.ErrNoCookie) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(permissionDenied))
			return
		}

		_, user, err := s.authService.Authcheck(r.Context(), session.Value)
		if err != nil || user == nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(permissionDenied))
			return
		}

		if !model.Roles(user.Role).IsValid() {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(permissionDenied))
			return
		}

		next.ServeHTTP(w, r)
	})
}
