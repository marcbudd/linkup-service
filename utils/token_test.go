package utils_test

import (
	"testing"

	"github.com/marcbudd/linkup-service/utils"
)

func TestGenerateTokenString(t *testing.T) {
	length := 32
	token, err := utils.GenerateTokenString(length)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	// Test length of token
	if len(token) != length {
		t.Errorf("Unexpected token length. Expected: %d, Got: %d", length, len(token))
	}
}
