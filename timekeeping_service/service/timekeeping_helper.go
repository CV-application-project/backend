package service

import (
	"Backend-Server/timekeeping_service/api"
	"Backend-Server/timekeeping_service/store"
	"encoding/json"
)

func (s *Service) convertToAPIHistory(histories []store.TimekeepingHistory) []*api.TimekeepingHistory {
	convertedList := make([]*api.TimekeepingHistory, 0, len(histories))
	for _, history := range histories {
		var data *api.HistoryDetail
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

func (s *Service) getOption(req *api.UpdateHistoryOfUserRequest) option {
	if req.StartTime == 0 && req.EndTime == 0 {
		return TypeUnspecified
	}
	if req.StartTime == 0 {
		return TypeUpdateEndTime
	}
	if req.EndTime == 0 {
		return TypeUpdateStartTime
	}
	return TypeUnspecified
}

func (s *Service) getHistoryData(history store.TimekeepingHistory) *api.HistoryDetail {
	var data api.HistoryDetail
	if err := json.Unmarshal([]byte(history.Data.String), &data); err != nil {
		return nil
	}
	return &data
}
