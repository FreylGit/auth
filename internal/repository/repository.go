package repository

import (
	"context"
	"github.com/FreylGit/auth/internal/model"
)

type UserRepository interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.User) (int64, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
	AddRole(ctx context.Context, userId int64, roleId int64) error
	RemoveRole(ctx context.Context, userId int64, roleId int64) error
}

type RoleRepository interface {
	Get(ctx context.Context, id int64) (*model.Role, error)
	GetByName(ctx context.Context, name string) (*model.Role, error)
	Create(ctx context.Context, role *model.Role) (int64, error)
	Update(ctx context.Context, role *model.Role) error
	Delete(ctx context.Context, id int64) error
}
