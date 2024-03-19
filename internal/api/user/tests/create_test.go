package tests

import (
	"context"
	"fmt"
	user2 "github.com/FreylGit/auth/internal/api/user"
	"github.com/FreylGit/auth/internal/model"
	"github.com/FreylGit/auth/internal/service"
	"github.com/FreylGit/auth/internal/service/mocks"
	desc "github.com/FreylGit/auth/pkg/user_v1"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}
	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id               = int64(2)
		name             = "andrey"
		email            = "email@ge.ty"
		password         = "123dver@f23"
		password_confirm = password
		roleReq          = desc.Role(0)

		serviceErr = fmt.Errorf("service erroro")
		_          = serviceErr
		req        = &desc.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password_confirm,
			Role:            roleReq,
		}

		role = model.Role{
			Id:   int64(roleReq),
			Name: "user",
		}
		user = &model.User{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     role,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)
	defer t.Cleanup(mc.Finish)
	tests := []struct {
		name                string
		args                args
		want                *desc.CreateResponse
		err                 error
		userServiceMockFunc userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMockFunc: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, user).Return(id, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userServiceMockFunc := tt.userServiceMockFunc(mc)
			api := user2.NewImplementation(userServiceMockFunc)
			createResp, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want.Id, createResp.Id)
		})
	}
}
