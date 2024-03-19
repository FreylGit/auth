package auth

import (
	"context"
	desc "github.com/FreylGit/auth/pkg/auth_v1"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	rtoken, err := i.authService.GetRefreshToken(ctx, req.GetOldRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: rtoken,
	}, nil
}
