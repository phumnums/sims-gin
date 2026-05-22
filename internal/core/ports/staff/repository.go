package user

import (
	"cims/internal/core/domain"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	GetRoleByID(ctx context.Context, roleID string) (*domain.Roles, error)
	GetCampusByID(ctx context.Context, campusID string) (*domain.Campus, error)
	CheckDuplicateUsername(ctx context.Context, username string) error
	CheckDuplicatePhoneNumber(ctx context.Context, phone_number string) error
	CheckDuplicateEmail(ctx context.Context, email string) error
	CreateUser(ctx context.Context, tx *gorm.DB, user domain.Users) error
	CreateStaff(ctx context.Context, tx *gorm.DB, staff domain.Staffs) error
}
