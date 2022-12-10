package service

import (
	"Backend-Server/timekeeping_service/api"
	"Backend-Server/timekeeping_service/store"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type SearchType int

const (
	ByMonth SearchType = 1
	ByYear  SearchType = 2
	ByDay   SearchType = 3
)

func (s *Service) GetHistoryOfUser(ctx context.Context, req *api.GetHistoryOfUserRequest) (*api.GetHistoryOfUserResponse, error) {
	logger := s.log.WithName("Service | GetHistoryOfUser").WithValues("traceId", req.UserId)
	var err error
	if err = req.Validate(); err != nil {
		logger.Error(err, "Validate")
		return nil, err
	}
	if int(req.Time.Year) > time.Now().Year() {
		logger.Info("Year is invalid")
		return nil, errors.New("year is invalid")
	}
	var timekeepingHistories []store.TimekeepingHistory
	switch findTypeOfTime(req.Time) {
	case ByYear:
		timekeepingHistories, err = s.store.GetTimekeepingHistoryInYearByUserId(ctx, store.GetTimekeepingHistoryInYearByUserIdParams{
			UserID: req.UserId,
			Year:   int32(req.Time.Year),
		})
		if err != nil {
			logger.Error(err, "Store | GetTimekeepingHistoryInYearByUserId")
			return nil, err
		}
	case ByMonth:
		timekeepingHistories, err = s.store.GetTimekeepingHistoryAtMonthByUserId(ctx, store.GetTimekeepingHistoryAtMonthByUserIdParams{
			UserID: req.UserId,
			Month:  int32(req.Time.Month),
			Year:   int32(req.Time.Year),
		})
		if err != nil {
			logger.Error(err, "Store | GetTimekeepingHistoryInYearByUserId | Can not get timekeeping in month")
			return nil, err
		}
	default:
		dayTimekeepingHistories, err := s.store.GetTimekeepingHistoryInDayByUserId(ctx, store.GetTimekeepingHistoryInDayByUserIdParams{
			UserID: req.UserId,
			Day:    int32(req.Time.Day),
			Month:  int32(req.Time.Month),
			Year:   int32(req.Time.Year),
		})
		if err != nil {
			logger.Error(err, "Store | GetTimekeepingHistoryInYearByUserId")
			return nil, err
		}
		timekeepingHistories = append(timekeepingHistories, dayTimekeepingHistories)
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
	if req.StartTime > req.EndTime {
		err = errors.New("invalid time")
		logger.Error(err, "Invalid time")
		return nil, err
	}
	if _, err = s.store.GetTimekeepingHistoryInDayByUserId(ctx, store.GetTimekeepingHistoryInDayByUserIdParams{
		UserID: req.UserId,
		Day:    int32(req.Time.Day),
		Month:  int32(req.Time.Month),
		Year:   int32(req.Time.Year),
	}); err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err, "GetTimekeepingHistoryInDayByUserId")
			return nil, err
		}
	}
	data := &api.Data{
		TotalTime: uint64(req.EndTime - req.StartTime),
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Details: []*api.Data_Line{
			{
				Index:     0,
				StartTime: req.StartTime,
				EndTime:   req.EndTime,
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
		Day:    int32(req.Time.Day),
		Month:  int32(req.Time.Month),
		Year:   int32(req.Time.Year),
		Data:   sql.NullString{Valid: true, String: string(dataBytes)},
	}); err != nil {
		logger.Error(err, "CreateTimekeepingHistory")
		return nil, err
	}
	return &api.CreateHistoryOfUserResponse{Code: http.StatusOK, Message: "success"}, nil
}

func (s *Service) convertToAPIHistory(histories []store.TimekeepingHistory) []*api.TimekeepingHistory {
	convertedList := make([]*api.TimekeepingHistory, 0, len(histories))
	for _, history := range histories {
		var data *api.Data
		if err := json.Unmarshal([]byte(history.Data.String), &data); err != nil {
			s.log.Error(err, "Can not unmarshal data")
			return nil
		}
		convertedList = append(convertedList, &api.TimekeepingHistory{
			UserId: history.UserID,
			Day:    uint32(history.Day),
			Month:  uint32(history.Month),
			Year:   uint32(history.Year),
			Data:   data,
		})
	}
	return convertedList
}

func findTypeOfTime(historyTime *api.HistoryTime) SearchType {
	if historyTime.Day == 0 {
		if historyTime.Month == 0 {
			return ByYear
		}
		return ByMonth
	}
	return ByDay
}
