package service

import (
	"encoding/json"
	"net/http"
)

type AppHandleFunc func(w http.ResponseWriter, r *http.Request) error
type AppMiddleware func(AppHandleFunc) AppHandleFunc

type AppHandler struct {
	handle      AppHandleFunc
	middlewares []AppMiddleware
}

func New(middlewares ...AppMiddleware) *AppHandler {
	return &AppHandler{middlewares: append(([]AppMiddleware)(nil), middlewares...)}
}

func (h *AppHandler) Build(handleFunc AppHandleFunc) *AppHandler {
	handler := &AppHandler{
		handle:      handleFunc,
		middlewares: h.middlewares,
	}
	for i := range handler.middlewares {
		handler.handle = handler.middlewares[len(handler.middlewares)-1-i](handler.handle)
	}
	return handler
}

func (h *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.handle == nil {
		return
	}
	err := h.handle(w, r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		message := struct {
			Message string
			Error   string
		}{
			Message: "failed",
			Error:   err.Error(),
		}
		_ = json.NewEncoder(w).Encode(message)
	}
}

func (s *Service) HTTPHandler(httpMux *http.ServeMux) {
	handler := New(s.authenticationMiddleware, s.corsMiddleware)
	httpMux.Handle("/cic/upsert", handler.Build(s.HTTPUpsertCICForUser))
	httpMux.Handle("/face/register", handler.Build(s.HTTPRegisterNewUserFace))
	httpMux.Handle("/face/authorize", handler.Build(s.HTTPAuthorizeNewUserFace))
	httpMux.Handle("/timekeeping/create", handler.Build(s.HTTPCreateHistoryOfUser))
	httpMux.Handle("/timekeeping/history", handler.Build(s.HTTPGetHistoryOfUser))
	httpMux.Handle("/timekeeping/update", handler.Build(s.HTTPUpdateHistoryOfUser))
	httpMux.Handle("/timekeeping/upsert", handler.Build(s.HTTPUpsertHistoryOfUser))
	httpMux.Handle("/user/update", handler.Build(s.HTTPUpdateUser))
	httpMux.Handle("/user/export", handler.Build(s.HTTPExportUsersByDepartment))
}
