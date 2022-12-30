package service

import (
	"Backend-Server/common/ctx_key"
	cvApi "Backend-Server/cv_service/api"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"io"
	"net/http"
)

func (s *Service) HTTPUpsertCICForUser(res http.ResponseWriter, req *http.Request) error {
	logger := s.log.WithName("HTTPUpsertCICForUser")
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
	userId := req.Context().Value(ctx_key.ContextUserId)
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
	resp, err := s.cvClient.UpsertCICForUser(context.Background(), &cvApi.UpsertCICForUserRequest{
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
