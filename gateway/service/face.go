package service

import (
	cvApi "Backend-Server/cv_service/api"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"strings"
)

func (s *Service) HTTPRegisterNewUserFace(writer http.ResponseWriter, req *http.Request) error {
	logger := s.log.WithName("HTTPRegisterNewUserFace")
	if req.Method != "POST" {
		err := errors.New("invalid HTTP method")
		logger.Error(err, "POST")
		return err
	}
	if err := req.ParseMultipartForm(100 << 20); err != nil {
		logger.Error(err, "ParseMultipartForm")
		return err
	}
	faceFile, header, err := req.FormFile("face")
	if err != nil {
		logger.Error(err, "FormFile")
		return err
	}
	defer faceFile.Close()
	userId := strings.Split(header.Filename, ".")[0]
	userIdInt := cast.ToInt64(userId)
	if userIdInt != req.Context().Value(ContextUserId) {
		err = errors.New("wrong user")
		return err
	}
	logger.WithValues("userId", userId)
	dataBytes := bytes.NewBuffer(nil)
	if _, err := io.Copy(dataBytes, faceFile); err != nil {
		logger.Error(err, "io.Copy")
		return err
	}
	res, err := s.cvClient.RegisterUserFace(context.Background(), &cvApi.RegisterUserFaceRequest{
		UserId: userIdInt,
		Image:  dataBytes.Bytes(),
	})
	if err != nil {
		logger.Error(err, "cvClient | RegisterUserFace")
		return err
	}
	if err = json.NewEncoder(writer).Encode(res); err != nil {
		logger.Error(err, "json | NewEncoder | Encode")
		return err
	}
	return nil
}

func (s *Service) HTTPAuthorizeNewUserFace(writer http.ResponseWriter, req *http.Request) error {
	logger := s.log.WithName("HTTPAuthorizeNewUserFace")
	if req.Method != "POST" {
		err := errors.New("invalid HTTP method")
		logger.Error(err, "POST")
		return err
	}
	if err := req.ParseMultipartForm(100 << 20); err != nil {
		logger.Error(err, "ParseMultipartForm")
		return err
	}
	faceFile, header, err := req.FormFile("face")
	if err != nil {
		logger.Error(err, "FormFile")
		return err
	}
	defer faceFile.Close()
	userId := strings.Split(header.Filename, ".")[0]
	userIdInt := cast.ToInt64(userId)
	if userIdInt != req.Context().Value(ContextUserId) {
		err = errors.New("wrong user")
		return err
	}
	logger.WithValues("userId", userId)
	dataBytes := bytes.NewBuffer(nil)
	if _, err := io.Copy(dataBytes, faceFile); err != nil {
		logger.Error(err, "io.Copy")
		return err
	}
	res, err := s.cvClient.AuthorizeUserFace(context.Background(), &cvApi.AuthorizeUserFaceRequest{
		UserId: userIdInt,
		Image:  dataBytes.Bytes(),
	})
	if err != nil {
		logger.Error(err, "cvClient | AuthorizeUserFace")
		return err
	}
	if err = json.NewEncoder(writer).Encode(res); err != nil {
		logger.Error(err, "json | NewEncoder | Encode")
		return err
	}
	return nil
}
