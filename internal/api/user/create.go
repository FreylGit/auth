package user

import (
	"context"
	"github.com/FreylGit/auth/internal/converter"
	desc "github.com/FreylGit/auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	user := converter.ToUserFromDesc(req.GetName(), req.GetEmail(), req.GetPassword(), req.GetRole())
	id, err := i.userService.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &desc.CreateResponse{Id: id}, nil

}
