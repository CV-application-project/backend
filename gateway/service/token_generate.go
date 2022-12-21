package service

import (
	"Backend-Server/user_service/api"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const secretKey string = "swt_very_long_secret_key_that_can_not_be_break"

type jwtCustomClaim struct {
	int64
	jwt.StandardClaims
}

func (j *jwtCustomClaim) Valid() error {
	if j.ExpiresAt < time.Now().UTC().Unix() {
		return errors.New("expired time is in the pass")
	}
	return nil
}

type BearerToken struct {
	Token     string
	ExpiredAt time.Time
}

func createBearerToken(tokenInfo *api.UserToken) (*BearerToken, error) {
	currentTime := time.Now().UTC()
	claims := &jwtCustomClaim{
		tokenInfo.UserId,
		jwt.StandardClaims{
			ExpiresAt: tokenInfo.ExpiredAt,
			Issuer:    "admin",
			IssuedAt:  currentTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	finalToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	return &BearerToken{
		Token:     finalToken,
		ExpiredAt: time.Unix(tokenInfo.ExpiredAt, 0),
	}, err
}
