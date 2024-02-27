package user

import (
	"github.com/FreylGit/auth/internal/client/db"
	"github.com/FreylGit/auth/internal/repository"
	"github.com/FreylGit/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	roleRepository repository.RoleRepository
	txManager      db.TxManager
}

func NewService(userRepository repository.UserRepository,
	roleRepository repository.RoleRepository,
	txManager db.TxManager) service.UserService {
	return &serv{
		userRepository: userRepository,
		roleRepository: roleRepository,
		txManager:      txManager,
	}
}
