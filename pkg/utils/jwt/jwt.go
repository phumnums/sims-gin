package jwt

import (
	"cims/internal/infarstructure/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	StaffID  string `json:"staff_id"`
	RoleName string `json:"role_name"`
	CumpusID string `json:"cumpus_id"`
	jwt.RegisteredClaims
}

func getSecretKey() []byte {
	return []byte(config.App.Server.Secret)
}

func GenerageToken(staffID, roleName, campusID string) (string, error) {
	claims := Claims{
		StaffID:  staffID,
		RoleName: roleName,
		CumpusID: campusID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(getSecretKey())
}
