package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateTokenString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	token := base64.RawURLEncoding.EncodeToString(randomBytes)
	if len(token) > length {
		token = token[:length]
	}

	return token, nil
}
