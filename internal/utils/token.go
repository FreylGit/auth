package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

type CustomClaims struct {
	jwt.StandardClaims
	RefreshId string `json:"refreshId"`
	Email     string `json:"email"`
	UserId    string `json:"userId"`
}

func GenerateAccessToken(userId string, email string, refreshId string, exp time.Time, secretKey []byte) (string, error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
		RefreshId: refreshId,
		Email:     email,
		UserId:    userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func GenerateRefreshToken(userId string, exp time.Time, secretKey []byte) (string, error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
		UserId: userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenStr string, secretKey []byte) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected token signing method")
			}

			return secretKey, nil
		},
	)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
