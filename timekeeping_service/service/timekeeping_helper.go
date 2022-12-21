package service

import (
	"Backend-Server/timekeeping_service/api"
	"Backend-Server/timekeeping_service/store"
	"encoding/json"
)

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
