package service

import (
	"Backend-Server/cv_service/store"
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"gocv.io/x/gocv"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func (s *Service) registerPreProcessing(image []byte) (string, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(rootPath, "cv_service/service/face_verification")
	filePath := filepath.Join(path, "face.jpg")
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := file.Write(image); err != nil {
		return "", err
	}

	faceImage := gocv.IMRead(filePath, gocv.IMReadColor)
	faceImage = removeNoise(faceImage)
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	if !classifier.Load("cv_service/models/haarcascade_frontalface_default.xml") {
		return "", fmt.Errorf("can't load cascade classifier")
	}
	faces := classifier.DetectMultiScale(faceImage)
	if len(faces) != 1 {
		return "", fmt.Errorf("picture must contain only 1 face")
	}
	cropImage := faceImage.Region(faces[0])
	if ok := gocv.IMWrite(filePath, cropImage); !ok {
		return "", fmt.Errorf("can not write face again to file")
	}
	return path, nil
}

func (s *Service) registerProcessing(path string) (string, error) {
	filename := filepath.Join(path, "face.jpg")
	cmd := exec.Command("python", filepath.Join(path, "main.py"), "--register", filename)
	output := bytes.NewBuffer(nil)
	cmd.Stdout = output
	if err := cmd.Run(); err != nil {
		return "", err
	}
	arr := strings.Split(output.String(), "\n")
	fmt.Print(arr)
	if len(arr) != 2 {
		return "", fmt.Errorf("output len invalid")
	}
	filename = arr[0]
	return filename, nil
}

func (s *Service) registerPostProcessing(ctx context.Context, userId int64, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		s.log.Error(err, "open file fail")
		return err
	}
	defer file.Close()
	dataBytes := bytes.NewBuffer(nil)
	if _, err = io.Copy(dataBytes, file); err != nil {
		s.log.Error(err, "copy file fail")
		return err
	}
	if _, err = s.store.CreateFaceByUserId(ctx, store.CreateFaceByUserIdParams{
		UserID: userId,
		Data:   sql.NullString{Valid: true, String: dataBytes.String()},
	}); err != nil {
		return err
	}
	if err = os.Remove(filename); err != nil {
		return err
	}
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(rootPath, "cv_service/service/face_verification")
	filePath := filepath.Join(path, "face.jpg")
	if err = os.Remove(filePath); err != nil {
		return err
	}
	return nil
}

func (s *Service) authorizePreProcessing(unknown []byte, known []byte) (string, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	rootPath = filepath.Join(rootPath, "cv_service/service/face_verification")
	unknownPath := filepath.Join(rootPath, "unknown.jpg")
	unknownFile, err := os.Create(unknownPath)
	if err != nil {
		return "", err
	}
	defer unknownFile.Close()

	if _, err := unknownFile.Write(unknown); err != nil {
		return "", err
	}

	faceImage := gocv.IMRead(unknownPath, gocv.IMReadColor)
	faceImage = removeNoise(faceImage)
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	if !classifier.Load("cv_service/models/haarcascade_frontalface_default.xml") {
		return "", fmt.Errorf("can't load cascade classifier")
	}
	faces := classifier.DetectMultiScale(faceImage)
	if len(faces) != 1 {
		return "", fmt.Errorf("picture must contain only 1 face")
	}
	cropImage := faceImage.Region(faces[0])
	if ok := gocv.IMWrite(unknownPath, cropImage); !ok {
		return "", fmt.Errorf("can not write face again to file")
	}

	knownFile, err := os.Create(filepath.Join(rootPath, "known.txt"))
	if err != nil {
		return "", err
	}
	defer knownFile.Close()

	if _, err = knownFile.Write(known); err != nil {
		return "", err
	}
	return rootPath, nil
}

func (s *Service) authorizeProcessing(path string) error {
	cmd := exec.Command("python", filepath.Join(path, "main.py"), "--authorize", filepath.Join(path, "unknown.jpg"), filepath.Join(path, "known.txt"))
	matchResult := bytes.NewBuffer(nil)
	cmd.Stdout = matchResult
	if err := cmd.Run(); err != nil {
		return err
	}
	if matchResult.String() == "true\n" {
		return nil
	}
	return fmt.Errorf("face not match")
}

func (s *Service) authorizePostProcessing(path string) error {
	unknownPath := filepath.Join(path, "unknown.jpg")
	if err := os.Remove(unknownPath); err != nil {
		return err
	}

	knownPath := filepath.Join(path, "known.txt")
	if err := os.Remove(knownPath); err != nil {
		return err
	}
	return nil
}
