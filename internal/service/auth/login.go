package auth

import (
	"context"
	"github.com/FreylGit/auth/internal/model"
	"github.com/FreylGit/auth/internal/utils"
	"github.com/pkg/errors"
	passUtils "github.com/vzglad-smerti/password_hash"
	"strconv"
	"time"
)

func (s *serv) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.userRepository.GetByEmail(ctx, email)
	isEquals, err := passUtils.Verify(user.Password, password)
	if err != nil {
		return "", err
	}
	if !isEquals {
		return "", errors.New("Incorrect password")
	}

	exp := time.Now().Add(time.Hour * 24)
	rtoken, err := utils.GenerateRefreshToken(strconv.FormatInt(user.Id, 10), exp, []byte("test"))
	if err != nil {
		return "", err
	}
	rtokenModel := &model.RefreshToken{
		Token: rtoken,
		Exp:   exp,
	}
	_, err = s.refreshTokenRepository.Create(ctx, rtokenModel)
	if err != nil {
		return "", err
	}
	return rtoken, nil
}
