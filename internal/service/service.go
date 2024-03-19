package service

import (
	"context"
	"github.com/FreylGit/auth/internal/model"
)

type UserService interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.User) (int64, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
}

type AuthService interface {
	Login(ctx context.Context, email string, password string) (string, error)
	GetRefreshToken(ctx context.Context, rToken string) (string, error)
	GetAccessToken(ctx context.Context, rToken string) (string, error)
}
