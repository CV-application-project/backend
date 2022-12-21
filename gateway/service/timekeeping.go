package service

import (
	timekeepingApi "Backend-Server/timekeeping_service/api"
	"encoding/json"
	"errors"
	"net/http"
)

func (s *Service) HTTPCreateHistoryOfUser(w http.ResponseWriter, r *http.Request) error {
	logger := s.log.WithName("HTTPCreateHistoryOfUser")
	var err error
	if r.Method != "POST" {
		err = errors.New("invalid HTTP method")
		logger.Error(err, "POST")
		return err
	}
	var request timekeepingApi.CreateHistoryOfUserRequest
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error(err, "json | decode")
		return err
	}
	if r.Context().Value(ContextUserId) != request.UserId {
		err = errors.New("wrong user")
		return err
	}
	resp, err := s.timekeepingClient.CreateHistoryOfUser(r.Context(), &request)
	if err != nil {
		logger.Error(err, "timekeepingClient | CreateHistoryOfUser")
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		logger.Error(err, "json | encode")
		return err
	}
	return nil
}

func (s *Service) HTTPGetHistoryOfUser(w http.ResponseWriter, r *http.Request) error {
	logger := s.log.WithName("HTTPCreateHistoryOfUser")
	var err error
	if r.Method != "POST" {
		err = errors.New("invalid HTTP method")
		logger.Error(err, "POST")
		return err
	}
	var request timekeepingApi.GetHistoryOfUserRequest
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error(err, "json | decode")
		return err
	}
	if r.Context().Value(ContextUserId) != request.UserId {
		err = errors.New("wrong user")
		return err
	}
	resp, err := s.timekeepingClient.GetHistoryOfUser(r.Context(), &request)
	if err != nil {
		logger.Error(err, "timekeepingClient | GetHistoryOfUser")
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		logger.Error(err, "json | encode")
		return err
	}
	return nil
}
func (s *Service) HTTPUpdateHistoryOfUser(w http.ResponseWriter, r *http.Request) error {
	logger := s.log.WithName("HTTPCreateHistoryOfUser")
	var err error
	if r.Method != "POST" {
		err = errors.New("invalid HTTP method")
		logger.Error(err, "POST")
		return err
	}
	var request timekeepingApi.UpdateHistoryOfUserRequest
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error(err, "json | decode")
		return err
	}
	if r.Context().Value(ContextUserId) != request.UserId {
		err = errors.New("wrong user")
		return err
	}
	resp, err := s.timekeepingClient.UpdateHistoryOfUser(r.Context(), &request)
	if err != nil {
		logger.Error(err, "timekeepingClient | UpdateHistoryOfUser")
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		logger.Error(err, "json | encode")
		return err
	}
	return nil
}
