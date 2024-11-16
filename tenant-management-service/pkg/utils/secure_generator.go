package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateSecureToken generates a cryptographically secure random token.
func GenerateSecureToken(length int) (string, error) {
	// Generate random bytes
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Encode to base64 to make it URL-safe
	return base64.URLEncoding.EncodeToString(bytes), nil
}
