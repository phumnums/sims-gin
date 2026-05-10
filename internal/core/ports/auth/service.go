package auth

import "cims/internal/adapters/dto"

type Service interface {
	Login(req dto.LoginRequest) (string, int, error)
}
