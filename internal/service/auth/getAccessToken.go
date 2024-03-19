package auth

import (
	"context"
	"github.com/FreylGit/auth/internal/utils"
	"strconv"
	"time"
)

func (s *serv) GetAccessToken(ctx context.Context, rToken string) (string, error) {
	tokenModel, err := s.refreshTokenRepository.Get(ctx, rToken)
	if err != nil {
		panic(err.Error())
	}
	exp := time.Now().Add(time.Hour * 2)
	claims, err := utils.VerifyToken(rToken, []byte("test"))
	if err != nil {
		panic(err.Error())
	}
	userId := claims.UserId
	userIdInt, err := strconv.ParseInt(userId, 10, 64)

	userModel, err := s.userRepository.Get(ctx, userIdInt)

	refreshId := strconv.FormatInt(tokenModel.Id, 10)

	newRToken, err := utils.GenerateAccessToken(userId, userModel.Email, refreshId, exp, []byte("test"))
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
