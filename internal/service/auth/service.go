package auth

import (
	"github.com/FreylGit/auth/internal/repository"
	"github.com/FreylGit/auth/internal/service"
	"github.com/FreylGit/platform_common/pkg/db"
)

type serv struct {
	txManager              db.TxManager
	userRepository         repository.UserRepository
	refreshTokenRepository repository.RefreshTokenRepository
}

func NewService(txManager db.TxManager, userRepository repository.UserRepository, refreshTokenRepository repository.RefreshTokenRepository) service.AuthService {
	return &serv{
		txManager:              txManager,
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
	}
}
