package service

import (
	"Backend-Server/cv_service/api"
	"Backend-Server/cv_service/store"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"gocv.io/x/gocv"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var imageList = []string{"front.jpg", "back.jpg"}

func (s *Service) RegisterCICForUser(ctx context.Context, req *api.RegisterCICForUserRequest) (*api.RegisterCICForUserResponse, error) {
	if err := req.Validate(); err != nil {
		s.log.Error(err, "Validate")
		return nil, err
	}
	logger := s.log.WithName("RegisterCICForUser").WithValues("user_id", req.UserId)

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
	if _, err = s.store.CreateCICByUserId(ctx, store.CreateCICByUserIdParams{
		UserID: req.UserId,
		Data:   sql.NullString{Valid: true, String: string(dataBytes)},
	}); err != nil {
		logger.Error(err, "store | CreateCICByUserId")
		return nil, err
	}
	return &api.RegisterCICForUserResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}

func (s *Service) cicPreProcessing(userId int64, front []byte, back []byte) error {
	logger := s.log.WithName("cicPreProcessing").WithValues("user_id", userId)
	// handle front & back imageFile

	frontImageFile, err := os.Create(imageList[0])
	if err != nil {
		logger.Error(err, "OS | Create")
		return err
	}
	defer frontImageFile.Close()
	if _, err = frontImageFile.Write(front); err != nil {
		logger.Error(err, "File | Write")
		return err
	}
	// Read imageFile
	image := gocv.IMRead(imageList[0], gocv.IMReadColor)

	// Remove noise
	image = removeNoise(image)

	// Rewrite imageFile to file
	gocv.IMWrite(imageList[1], image)

	backImageFile, err := os.Create(imageList[1])
	if err != nil {
		logger.Error(err, "OS | Create")
		return err
	}
	defer backImageFile.Close()
	if _, err = backImageFile.Write(back); err != nil {
		logger.Error(err, "File | Write")
		return err
	}
	// Read imageFile
	image = gocv.IMRead(imageList[1], gocv.IMReadColor)

	// Remove noise
	image = removeNoise(image)

	// Rewrite imageFile to file
	gocv.IMWrite(imageList[1], image)

	return nil
}

func removeNoise(image gocv.Mat) gocv.Mat {
	gocv.MedianBlur(image, &image, 5)
	return image
}

func (s *Service) cicProcessing() ([]string, error) {
	finalText := make([]string, 0)
	for _, imageFile := range imageList {
		cmd := exec.Command("tesseract", imageFile, "stdout", "-l", "vie")
		var output bytes.Buffer
		cmd.Stdout = &output
		if err := cmd.Run(); err != nil {
			return nil, err
		}
		texts := strings.Split(output.String(), "\n")
		finalText = append(finalText, texts...)
	}
	return finalText, nil
}

func (s *Service) cicPostProcessing(texts []string) (*api.CitizenIdentityCard, error) {
	card := &api.CitizenIdentityCard{}
	for index, text := range texts {
		if strings.Contains(text, "Có giá trị đến") {
			expireDataString := strings.Split(text, ": ")[1]
			expireTime, err := convertToTime(expireDataString)
			if err != nil {
				return nil, err
			}
			card.ExpireDate = expireTime
		} else if strings.Contains(text, "số") {
			card.Id = strings.Split(text, " ")[len(strings.Split(text, " "))-1]
		} else if strings.Contains(text, "Họ và tên") {
			card.Name = strings.Split(text, ": ")[len(strings.Split(text, ": "))-1]
		} else if strings.Contains(text, "Ngày, tháng, năm sinh") {
			birthdayString := strings.Split(text, ": ")[len(strings.Split(text, ": "))-1]
			birthdayTime, err := convertToTime(birthdayString)
			if err != nil {
				return nil, err
			}
			card.Birthday = birthdayTime
		} else if strings.Contains(text, "Giới tính") {
			arr := strings.Split(text, ": ")
			if strings.Split(arr[1], " ")[0] == "Nam" {
				card.Gender = api.Gender_MALE
			} else {
				card.Gender = api.Gender_FEMALE
			}
			card.Country = arr[2]
		} else if strings.Contains(text, "Quê quán") {
			card.BornProvince = strings.Split(text, ": ")[1]
		} else if strings.Contains(text, "Nơi thường trú") {
			temp := texts[index+3]
			if len(texts[index+3]) < 2 {
				temp = texts[index+4]
			}
			card.Location = strings.TrimSpace(strings.Join([]string{texts[index+2], temp}, ", "))
		} else if strings.Contains(text, "Đặc điểm nhân dạng") {
			card.Provider = strings.TrimSpace(strings.Join([]string{texts[index-2], texts[index-1]}, " "))
		}
	}
	for _, imageFile := range imageList {
		if err := os.Remove(imageFile); err != nil {
			return nil, err
		}
	}
	return card, nil
}

func convertToTime(timeString string) (int64, error) {
	// Default format = dd/mm/yyyy
	// Want: yyy-mm-dd
	array := strings.Split(timeString, "/")
	var finalTime string
	for i := range array {
		if finalTime == "" {
			finalTime = array[i]
		} else {
			finalTime = strings.Join([]string{finalTime, array[i]}, "-")
		}
	}
	finalTimeTime, err := time.Parse("02-01-2006", finalTime)
	if err != nil {
		return 0, err
	}
	return finalTimeTime.Unix(), nil
}