package middleware

import (
	jwtpkg "cims/pkg/utils/jwt"
	"cims/pkg/utils/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			response.Error(c, http.StatusUnauthorized, errors.New("missing token"))
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &jwtpkg.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret-key"), nil
		},
		)

		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, errors.New("invalid token"))
			c.Abort()
			return
		}

		claims := token.Claims.(*jwtpkg.Claims)

		c.Set("staff_id", claims.StaffID)
		c.Set("role_name", claims.RoleName)
		c.Set("campus_id", claims.CumpusID)

		c.Next()
	}
}
