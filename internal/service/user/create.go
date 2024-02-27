package user

import (
	"context"
	"github.com/FreylGit/auth/internal/model"
	password "github.com/vzglad-smerti/password_hash"
)

func (s serv) Create(ctx context.Context, user *model.User) (int64, error) {
	var userId int64
	var role *model.Role
	passwordHash, err := password.Hash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = passwordHash

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		role, errTx = s.roleRepository.GetByName(ctx, user.Role.Name)
		if errTx != nil {
			return errTx
		}
		userId, errTx = s.userRepository.Create(ctx, user)
		if errTx != nil {
			return errTx
		}
		errTx = s.userRepository.AddRole(ctx, userId, role.Id)
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return userId, nil
}
