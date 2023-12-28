package secretgen

import (
	"crypto/rand"
	"fmt"
)

// GenerateRandomSecret generates a random string with characters that
// are only in the given character set.
func GenerateRandomAlphaNumeric(n int) (string, error) {
	r, err := GenerateRandomSecret(n, []byte(CharSetAlphaNumeric))
	if err != nil {
		return "", err
	}

	return r, nil
}

// GenerateRandomSecret generates a random string with characters that
// are only in the given character set.
func GenerateRandomSecret(n int, charset []byte) (string, error) {
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("generate random secret: %w", err)
	}

	t := len(charset)

	if t > 0 {
		for i := 0; i < n; i++ {
			bytes[i] = charset[bytes[i]%byte(t)]
		}
	}

	return string(bytes), nil
}
