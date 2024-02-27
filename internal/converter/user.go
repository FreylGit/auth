package converter

import (
	"github.com/FreylGit/auth/internal/model"
	desc "github.com/FreylGit/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *model.User) *desc.GetResponse {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)

	}

	return &desc.GetResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserFromDesc(name string, email string, password string, role desc.Role) *model.User {
	return &model.User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     ToRoleFromDesc(role),
	}
}

func ToRoleFromDesc(role desc.Role) model.Role {
	if role == 1 {
		return model.Role{Name: "admin"}
	}

	return model.Role{Name: "user"}
}
