package auth

import (
	"context"
	"github.com/FreylGit/auth/internal/utils"
	"time"
)

func (s *serv) GetRefreshToken(ctx context.Context, rToken string) (string, error) {
	tokenModel, err := s.refreshTokenRepository.Get(ctx, rToken)
	if err != nil {
		panic(err.Error())
	}
	exp := time.Now().Add(time.Hour * 25)
	claims, err := utils.VerifyToken(rToken, []byte("test"))
	if err != nil {
		panic(err.Error())
	}
	userId := claims.UserId
	newRToken, err := utils.GenerateRefreshToken(userId, exp, []byte("test"))
	if err != nil {
		panic(err.Error())
	}
	tokenModel.Token = newRToken
	tokenModel.Exp = exp
	err = s.refreshTokenRepository.Update(ctx, tokenModel)
	if err != nil {
		panic(err.Error())
	}

	return newRToken, nil
}
