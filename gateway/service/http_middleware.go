package service

import (
	"Backend-Server/common/ctx_key"
	"Backend-Server/common/errorz"
	"context"
	"errors"
	"net/http"
	"strings"
)

func (s *Service) authenticationMiddleware(next AppHandleFunc) AppHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := context.Background()
		tokenKey := r.Header.Get("authorization")
		if tokenKey == "" {
			return errors.New("missing authorization token")
		}
		if !strings.Contains(tokenKey, "Bearer") {
			return errors.New("invalid authorization token")
		}
		tokenKey = strings.Split(tokenKey, " ")[1]
		info, err := s.store.GetUserInfoByToken(ctx, tokenKey)
		if err != nil {
			return errorz.ErrUserNotFound
		}
		ctx = context.WithValue(ctx, ctx_key.ContextUserId, info.UserID)
		ctx = context.WithValue(ctx, ctx_key.ContextRole, info.Role.String)
		ctx = context.WithValue(ctx, ctx_key.ContextDepartment, info.Department.String)
		r = r.WithContext(ctx)
		return next(w, r)
	}
}

func (s *Service) corsMiddleware(next AppHandleFunc) AppHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		if r.Context().Value(ctx_key.ContextUserId) == nil {
			return errors.New("missing user id in context")
		}
		return next(w, r)
	}
}
