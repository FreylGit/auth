package user

import (
	"github.com/FreylGit/auth/internal/service"
	desc "github.com/FreylGit/auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{userService: userService}
}
