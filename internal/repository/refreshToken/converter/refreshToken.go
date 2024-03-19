package converter

import (
	"github.com/FreylGit/auth/internal/model"
	modelRepo "github.com/FreylGit/auth/internal/repository/refreshToken/model"
)

func ToRefreshTokenFromRepo(token modelRepo.RefreshToken) *model.RefreshToken {
	return &model.RefreshToken{
		Id:    token.Id,
		Token: string(token.Token),
		Exp:   token.Exp,
	}
}
