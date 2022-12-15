package service

import (
	"errors"
	"net/http"
)

func (s *Service) authenticationMiddleware(next AppHandleFunc) AppHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r.Header.Get("authorization") == "" {
			return errors.New("authorization token is invalid")
		}
		return next(w, r)
	}
}

func (s *Service) corsMiddleware(next AppHandleFunc) AppHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		return next(w, r)
	}
}
