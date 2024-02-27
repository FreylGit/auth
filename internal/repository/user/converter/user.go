package converter

import (
	"github.com/FreylGit/auth/internal/model"
	model2 "github.com/FreylGit/auth/internal/repository/role/model"
	modelRepo "github.com/FreylGit/auth/internal/repository/user/model"
)

func ToUserFromRepo(user modelRepo.User) *model.User {
	return &model.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToRoleFromRepo(role model2.Role) model.Role {
	return model.Role{
		Id:   role.Id,
		Name: role.NameLower,
	}
}
