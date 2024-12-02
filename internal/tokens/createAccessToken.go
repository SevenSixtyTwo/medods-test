package tokens

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	ip string
	jwt.RegisteredClaims
}

func CreateAccessToken(secret, ip string) (string, error) {
	claims := &JwtCustomClaims{
		ip,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("token signing %v", err)
	}

	return t, nil
}
