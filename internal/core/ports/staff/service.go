package user

import (
	"cims/internal/adapters/dto"
	"cims/internal/adapters/middleware"
	"context"
)

type Service interface {
	Create(ctx context.Context, auth middleware.AuthContext, req dto.CreateStaffRequest) (int, error)
}
