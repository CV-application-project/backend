package service

import (
	cvApi "Backend-Server/cv_service/api"
	"Backend-Server/gateway/api"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"io"
	"net/http"
)

func (s *Service) RegisterCICForUser(ctx context.Context, req *api.RegisterCICForUserRequest) (*api.RegisterCICForUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	logger := s.log.WithName("RegisterCICForUser").WithValues("user_id", req.UserId)
	if req.Front == nil || req.Back == nil {
		logger.Info("image is empty")
		return nil, nil
	}
	_, err := s.cvClient.RegisterCICForUser(ctx, &cvApi.RegisterCICForUserRequest{
		UserId: req.UserId,
		Front:  req.Front,
		Back:   req.Back,
	})
	if err != nil {
		logger.Error(err, "cvClient | RegisterCICForUser")
		return nil, err
	}
	return &api.RegisterCICForUserResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}

func (s *Service) HTTPRegisterCICForUser(res http.ResponseWriter, req *http.Request) error {
	logger := s.log.WithName("HTTPRegisterCICForUser")
	if req.Method != "POST" {
		err := errors.New("invalid HTTP method")
		logger.Error(err, "POST")
		return err
	}
	if err := req.ParseMultipartForm(100 << 20); err != nil {
		logger.Error(err, "ParseMultipartForm")
		return err
	}
	frontImageFile, _, err := req.FormFile("front")
	if err != nil {
		logger.Error(err, "FormFile")
		return err
	}
	defer frontImageFile.Close()
	userId := req.Context().Value(ContextUserId)
	logger.WithValues("user_id", userId)
	frontData := bytes.NewBuffer(nil)
	if _, err = io.Copy(frontData, frontImageFile); err != nil {
		logger.Error(err, "Copy")
		return err
	}
	backImageFile, _, err := req.FormFile("back")
	if err != nil {
		logger.Error(err, "FormFile")
		return err
	}
	defer backImageFile.Close()
	backData := bytes.NewBuffer(nil)
	if _, err = io.Copy(backData, backImageFile); err != nil {
		logger.Error(err, "Copy")
		return err
	}
	resp, err := s.RegisterCICForUser(context.Background(), &api.RegisterCICForUserRequest{
		UserId: cast.ToInt64(userId),
		Front:  frontData.Bytes(),
		Back:   backData.Bytes(),
	})
	if err != nil {
		logger.Error(err, "RegisterCICForUser")
		return err
	}
	if err = json.NewEncoder(res).Encode(resp); err != nil {
		logger.Error(err, "Send response")
		return err
	}
	return nil
}
