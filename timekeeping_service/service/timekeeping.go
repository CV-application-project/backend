package service

import (
	"Backend-Server/timekeeping_service/api"
	"Backend-Server/timekeeping_service/store"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"math"
	"net/http"
	"time"
)

type option string

const (
	TypeUpdateEndTime   option = "end"
	TypeUpdateStartTime option = "start"
	TypeUnspecified     option = "unspecified"
)

func (s *Service) GetHistoryOfUser(ctx context.Context, req *api.GetHistoryOfUserRequest) (*api.GetHistoryOfUserResponse, error) {
	logger := s.log.WithName("Service | GetHistoryOfUser").WithValues("traceId", req.UserId)
	if req.Time.From > req.Time.To {
		logger.Info("invalid duration")
		return nil, errors.New("invalid duration")
	}
	timekeepingHistories, err := s.store.GetTimekeepingHistoryByDuration(ctx, store.GetTimekeepingHistoryByDurationParams{
		UserID:      req.UserId,
		CreatedAt:   time.Unix(req.Time.From, 0),
		CreatedAt_2: time.Unix(req.Time.To, 0),
	})
	if err != nil {
		logger.Error(err, "Store | GetTimekeepingHistoryByDuration")
		return nil, err
	}

	return &api.GetHistoryOfUserResponse{
		Data: s.convertToAPIHistory(timekeepingHistories),
	}, nil
}

func (s *Service) CreateHistoryOfUser(ctx context.Context, req *api.CreateHistoryOfUserRequest) (*api.CreateHistoryOfUserResponse, error) {
	logger := s.log.WithName("Service | CreateHistoryOfUser").WithValues("traceId", req.UserId)
	var err error
	if err = req.Validate(); err != nil {
		logger.Error(err, "Validate")
		return nil, err
	}
	now := time.Now()
	if _, err = s.store.GetTimekeepingHistoryInDayByUserId(ctx, store.GetTimekeepingHistoryInDayByUserIdParams{
		UserID: req.UserId,
		Day:    int32(now.Day()),
		Month:  int32(now.Month()),
		Year:   int32(now.Year()),
	}); err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err, "GetTimekeepingHistoryInDayByUserId")
			return nil, err
		}
	}
	data := &api.HistoryDetail{
		TotalTime: 0,
		StartTime: req.StartTime,
		Details: []*api.HistoryDetail_Line{
			{
				StartTime: req.StartTime,
			},
		},
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		logger.Error(err, "Marshal")
		return nil, err
	}
	if _, err = s.store.CreateTimekeepingHistory(ctx, store.CreateTimekeepingHistoryParams{
		UserID: req.UserId,
		Day:    int32(now.Day()),
		Month:  int32(now.Month()),
		Year:   int32(now.Year()),
		Data:   sql.NullString{Valid: true, String: string(dataBytes)},
	}); err != nil {
		logger.Error(err, "CreateTimekeepingHistory")
		return nil, err
	}
	return &api.CreateHistoryOfUserResponse{Code: http.StatusOK, Message: "success"}, nil
}

func (s *Service) UpdateHistoryOfUser(ctx context.Context, req *api.UpdateHistoryOfUserRequest) (*api.UpdateHistoryOfUserResponse, error) {
	logger := s.log.WithName("UpdateHistoryOfUser").WithValues("userId", req.UserId)
	updateOption := s.getOption(req)
	if updateOption == TypeUnspecified {
		err := errors.New("invalid time")
		logger.Error(err, "getOption")
		return nil, err
	}
	now := time.Unix(cast.ToInt64(math.Abs(cast.ToFloat64(req.StartTime-req.EndTime))), 0)
	history, err := s.store.GetTimekeepingHistoryInDayByUserId(ctx, store.GetTimekeepingHistoryInDayByUserIdParams{
		UserID: req.UserId,
		Day:    int32(now.Day()),
		Month:  int32(now.Month()),
		Year:   int32(now.Year()),
	})
	if err != nil {
		logger.Error(err, "Store | GetTimekeepingHistoryInDayByUserId")
		return nil, err
	}
	historyData := s.getHistoryData(history)
	if historyData == nil {
		logger.Info("getHistoryData")
		return nil, errors.New("can not parse history data")
	}
	switch updateOption {
	case TypeUpdateEndTime:
		nearestStartTime := historyData.Details[len(historyData.Details)-1].StartTime
		historyData.Details[len(historyData.Details)-1].EndTime = req.EndTime
		historyData.TotalTime += uint64(req.EndTime - nearestStartTime)
		historyData.EndTime = req.EndTime
	case TypeUpdateStartTime:
		historyData.Details = append(historyData.Details, &api.HistoryDetail_Line{
			StartTime: req.StartTime,
			EndTime:   0,
		})
	default:
		break
	}

	dataBytes, err := json.Marshal(historyData)
	if err != nil {
		logger.Error(err, "Marshal")
		return nil, err
	}
	if _, err = s.store.UpdateTimekeepingHistoryInDay(ctx, store.UpdateTimekeepingHistoryInDayParams{
		UserID: req.UserId,
		Day:    int32(now.Day()),
		Month:  int32(now.Month()),
		Year:   int32(now.Year()),
		Data:   sql.NullString{Valid: true, String: string(dataBytes)},
	}); err != nil {
		logger.Error(err, "Store | UpdateTimekeepingHistoryInDay")
		return nil, err
	}
	return &api.UpdateHistoryOfUserResponse{
		Code:      http.StatusOK,
		Message:   "success",
		TotalTime: int64(historyData.TotalTime),
	}, nil
}

func (s *Service) UpsertHistoryOfUser(ctx context.Context, req *api.UpsertHistoryOfUserRequest) (*api.UpsertHistoryOfUserResponse, error) {
	lockTime := time.Unix(req.LockTime, 0)
	history, err := s.store.GetTimekeepingHistoryInDayByUserId(ctx, store.GetTimekeepingHistoryInDayByUserIdParams{
		UserID: req.UserId,
		Day:    int32(lockTime.Day()),
		Month:  int32(lockTime.Month()),
		Year:   int32(lockTime.Year()),
	})
	switch err {
	case nil:
		// Update
		if !history.Mode.Bool {
			result, err := s.UpdateHistoryOfUser(ctx, &api.UpdateHistoryOfUserRequest{
				UserId:  req.UserId,
				EndTime: req.LockTime,
			})
			if err != nil {
				return nil, err
			}
			return &api.UpsertHistoryOfUserResponse{
				Code:      http.StatusOK,
				Message:   "success",
				TotalTime: result.TotalTime,
			}, nil
		}
		result, err := s.UpdateHistoryOfUser(ctx, &api.UpdateHistoryOfUserRequest{
			UserId:    req.UserId,
			StartTime: req.LockTime,
		})
		if err != nil {
			return nil, err
		}
		return &api.UpsertHistoryOfUserResponse{
			Code:      http.StatusOK,
			Message:   "success",
			TotalTime: result.TotalTime,
		}, nil
	case sql.ErrNoRows:
		_, err := s.CreateHistoryOfUser(ctx, &api.CreateHistoryOfUserRequest{
			UserId:    req.UserId,
			StartTime: req.LockTime,
		})
		if err != nil {
			return nil, err
		}
		return &api.UpsertHistoryOfUserResponse{
			Code:      http.StatusOK,
			Message:   "success",
			TotalTime: 0,
		}, nil
	default:
		return nil, err
	}
}
