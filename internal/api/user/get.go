package user

import (
	"context"
	desc "github.com/FreylGit/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt.Time),
	}, nil
}
