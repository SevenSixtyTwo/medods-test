package tokens

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func CreateRefreshToken() (string, []byte, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", nil, err
	}

	refreshToken := base64.StdEncoding.EncodeToString(b)

	refreshTokenHashed, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return "", nil, err
	}

	return refreshToken, refreshTokenHashed, nil
}
