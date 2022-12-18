package service

import (
	"Backend-Server/cv_service/api"
	"Backend-Server/cv_service/store"
	"context"
	"database/sql"
	"fmt"
	"github.com/Kagami/go-face"
	"github.com/spf13/cast"
	"gocv.io/x/gocv"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
	if err = s.facePreProcessing(req.Image); err != nil {
		logger.Error(err, "facePreProcessing")
		return nil, err
	}
	descriptor, err := s.faceProcessing()
	if err != nil {
		logger.Error(err, "faceProcessing")
		return nil, err
	}
	if err = s.facePostProcessing(ctx, req.UserId, descriptor); err != nil {
		logger.Error(err, "facePostProcessing")
		return nil, err
	}
	return &api.RegisterUserFaceResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}

func (s *Service) facePreProcessing(image []byte) error {
	file, err := os.Create("face.jpg")
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(image); err != nil {
		return err
	}

	faceImage := gocv.IMRead("image.jpg", gocv.IMReadColor)
	faceImage = removeNoise(faceImage)
	if ok := gocv.IMWrite("image.jpg", faceImage); !ok {
		return fmt.Errorf("can not write face again to file")
	}
	return nil
}

func (s *Service) faceProcessing() (*face.Descriptor, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	rec, err := face.NewRecognizer(filepath.Join(path, "../models"))
	if err != nil {
		return nil, err
	}
	defer rec.Close()
	faces, err := rec.RecognizeFile("face.jpg")
	if err != nil {
		return nil, err
	}
	if len(faces) != 1 {
		return nil, fmt.Errorf("picture must contain only 1 face")
	}

	return &faces[0].Descriptor, nil
}

func (s *Service) facePostProcessing(ctx context.Context, userId int64, descriptor *face.Descriptor) error {
	descriptorString, err := convertDescriptorToString(descriptor)
	if err != nil {
		return err
	}
	if _, err = s.store.CreateFaceByUserId(ctx, store.CreateFaceByUserIdParams{
		UserID: userId,
		Data:   sql.NullString{Valid: true, String: descriptorString},
	}); err != nil {
		return err
	}
	return nil
}

func convertDescriptorToString(descriptor *face.Descriptor) (string, error) {
	if descriptor == nil {
		return "", fmt.Errorf("descriptor is nil")
	}
	descriptorString := ""
	for _, description := range descriptor {
		if descriptorString == "" {
			descriptorString = fmt.Sprintf("%f", description)
		} else {
			descriptorString = strings.Join([]string{descriptorString, fmt.Sprintf("%f", description)}, ", ")
		}
	}
	return descriptorString, nil
}

func (s *Service) AuthorizeUserFace(ctx context.Context, req *api.AuthorizeUserFaceRequest) (*api.AuthorizeUserFaceResponse, error) {
	logger := s.log.WithName("AuthorizeUserFace").WithValues("userId", req.UserId)
	user, err := s.store.GetFaceByUserId(ctx, req.UserId)
	if err != nil {
		logger.Error(err, "Store | GetFaceByUserId")
		return nil, err
	}
	if !user.Data.Valid {
		logger.Info("User data not valid")
		return nil, nil
	}
	descriptor, err := getDescriptorFromString(user.Data.String)
	if err != nil {
		logger.Error(err, "getDescriptorFromString")
		return nil, err
	}
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	rec, err := face.NewRecognizer(filepath.Join(path, "../models"))
	if err != nil {
		return nil, err
	}
	defer rec.Close()
	rec.SetSamples([]face.Descriptor{descriptor}, []int32{int32(req.UserId)})
	recognizeFace, err := rec.Recognize(req.Image)
	if err != nil {
		logger.Error(err, "Recognize")
		return nil, err
	}
	if recognizeFace == nil {
		logger.Info("face not found")
		return nil, err
	}
	if len(recognizeFace) > 1 {
		return nil, fmt.Errorf("picture must contain only 1 face")
	}
	actualUserId := rec.Classify(recognizeFace[0].Descriptor)
	if actualUserId < 0 {
		return nil, fmt.Errorf("user id for face not found")
	}
	if actualUserId != int(req.UserId) {
		return nil, fmt.Errorf("user not match")
	}
	return &api.AuthorizeUserFaceResponse{
		Code:    http.StatusOK,
		Message: "success",
		UserId:  int64(actualUserId),
	}, nil
}

func getDescriptorFromString(str string) (face.Descriptor, error) {
	if len(str) == 0 {
		return face.Descriptor{}, fmt.Errorf("desctiptor string len invalid")
	}
	descriptor := face.Descriptor{}
	arr := strings.Split(str, ", ")
	for index := range arr {
		descriptor[index] = cast.ToFloat32(arr[index])
	}
	return descriptor, nil
}
