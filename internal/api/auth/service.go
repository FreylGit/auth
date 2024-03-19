package auth

import (
	"github.com/FreylGit/auth/internal/service"
	desc "github.com/FreylGit/auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{authService: authService}
}
