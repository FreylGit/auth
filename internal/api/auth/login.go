package auth

import (
	"context"
	desc "github.com/FreylGit/auth/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	rtoken, err := i.authService.Login(ctx, req.GetEmail(), req.GetPassword())

	return &desc.LoginResponse{
		RefreshToken: rtoken,
	}, err
}
