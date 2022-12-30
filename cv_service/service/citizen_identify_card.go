package service

import (
	"Backend-Server/cv_service/api"
	"Backend-Server/cv_service/store"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
)

var imageList = []string{"front.jpg", "back.jpg"}

type StoreOption string

const (
	StoreOptionUpdate StoreOption = "update"
	StoreOptionCreate StoreOption = "insert"
)

func (s *Service) GetCICByUserId(ctx context.Context, req *api.GetCICByUserIdRequest) (*api.GetCICByUserIdResponse, error) {
	logger := s.log.WithName("GetCICByUserId").WithValues("userId", req.UserId)
	cic, err := s.store.GetCICByUserId(ctx, req.UserId)
	if err != nil {
		logger.Error(err, "Store | GetCICByUserId")
		return nil, err
	}
	var card api.CitizenIdentityCard
	if err = json.Unmarshal([]byte(cic.Data.String), &card); err != nil {
		logger.Error(err, "json | Unmarshal")
		return nil, err
	}

	return &api.GetCICByUserIdResponse{
		UserId: req.UserId,
		Card:   &card,
	}, nil
}

func (s *Service) UpsertCICForUser(ctx context.Context, req *api.UpsertCICForUserRequest) (*api.UpsertCICForUserResponse, error) {
	logger := s.log.WithName("RegisterCICForUser").WithValues("user_id", req.UserId)
	var opt StoreOption
	_, err := s.store.GetCICByUserId(ctx, req.UserId)
	switch err {
	case nil:
		// Update
		opt = StoreOptionUpdate
		logger.Info("Update card")
	case sql.ErrNoRows:
		opt = StoreOptionCreate
		logger.Info("User hasn't registered card yet")
	default:
		logger.Error(err, "store | GetCICByUserId")
		return nil, err
	}
	if err := s.cicPreProcessing(req.UserId, req.Front, req.Back); err != nil {
		logger.Error(err, "cicPreProcessing")
		return nil, err
	}
	texts, err := s.cicProcessing()
	if err != nil {
		logger.Error(err, "cicProcessing")
		return nil, err
	}
	card, err := s.cicPostProcessing(texts)
	if err != nil {
		logger.Error(err, "cicPostProcessing")
		return nil, err
	}
	dataBytes, err := json.Marshal(card)
	if err != nil {
		logger.Error(err, "json | Marshal")
		return nil, err
	}
	switch opt {
	case StoreOptionUpdate:
		if _, err = s.store.UpdateCICByUserId(ctx, store.UpdateCICByUserIdParams{
			UserID: req.UserId,
			Data:   sql.NullString{Valid: true, String: string(dataBytes)},
		}); err != nil {
			logger.Error(err, "store | UpdateCICByUserId")
		}
	case StoreOptionCreate:
		if _, err = s.store.CreateCICByUserId(ctx, store.CreateCICByUserIdParams{
			UserID: req.UserId,
			Data:   sql.NullString{Valid: true, String: string(dataBytes)},
		}); err != nil {
			logger.Error(err, "store | CreateCICByUserId")
			return nil, err
		}
	default:
		logger.Info("Opt not in list")
		break
	}

	return &api.UpsertCICForUserResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    card,
	}, nil
}
