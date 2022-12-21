package service

import (
	"Backend-Server/cv_service/api"
	"context"
	"database/sql"
	"net/http"
)

func (s *Service) RegisterUserFace(ctx context.Context, req *api.RegisterUserFaceRequest) (*api.RegisterUserFaceResponse, error) {
	logger := s.log.WithName("RegisterUserFace").WithValues("userId", req.UserId)
	// Check if user already been registered
	_, err := s.store.GetFaceByUserId(ctx, req.UserId)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err, "Store | GetFaceByUserId")
			return nil, err
		}
	}
	path, err := s.registerPreProcessing(req.Image)
	if err != nil {
		logger.Error(err, "registerPreProcessing")
		return nil, err
	}
	filename, err := s.registerProcessing(path)
	if err != nil {
		logger.Error(err, "registerProcessing")
		return nil, err
	}
	if err = s.registerPostProcessing(ctx, req.UserId, filename); err != nil {
		logger.Error(err, "registerPostProcessing")
		return nil, err
	}
	return &api.RegisterUserFaceResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}

func (s *Service) AuthorizeUserFace(ctx context.Context, req *api.AuthorizeUserFaceRequest) (*api.AuthorizeUserFaceResponse, error) {
	logger := s.log.WithName("AuthorizeUserFace").WithValues("userId", req.UserId)
	user, err := s.store.GetFaceByUserId(ctx, req.UserId)
	if err != nil {
		logger.Error(err, "Store | GetFaceByUserId")
		return nil, err
	}
	path, err := s.authorizePreProcessing(req.Image, []byte(user.Data.String))
	if err != nil {
		logger.Error(err, "authorizePreProcessing")
		return nil, err
	}

	if err = s.authorizeProcessing(path); err != nil {
		logger.Error(err, "authorizeProcessing")
		return nil, err
	}

	if err = s.authorizePostProcessing(path); err != nil {
		return nil, err
	}
	return &api.AuthorizeUserFaceResponse{
		Code:    http.StatusOK,
		Message: "success",
		UserId:  req.UserId,
	}, nil
}
