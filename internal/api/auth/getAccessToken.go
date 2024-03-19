package auth

import (
	"context"
	desc "github.com/FreylGit/auth/pkg/auth_v1"
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	atoken, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetAccessTokenResponse{AccessToken: atoken}, nil
}
