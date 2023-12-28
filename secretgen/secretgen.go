package secretgen

import (
	"crypto/rand"
	"fmt"
)

// GenerateRandomAlphaNumeric generates a random alphanumeric string of
// length n using the CharSetAlphaNumeric.
// It returns the generated string and any error encountered during generation.
func GenerateRandomAlphaNumeric(n int) (string, error) {
	r, err := GenerateRandomSecret(n, CharSetAlphaNumeric)
	if err != nil {
		return "", fmt.Errorf("generate alphanumeric: %w", err)
	}

	return r, nil
}

// GenerateRandomSecret generates a random secret string of the specified
// length using the provided character set. Returns an error if generation
// fails or if the inputs are invalid.
func GenerateRandomSecret(n int, charset string) (string, error) {
	if n <= 0 {
		return "",
			fmt.Errorf("invalid length, n must be greater than 0")
	}

	if len(charset) == 0 {
		return "",
			fmt.Errorf("empty character set, charset is required")
	}

	// Generate random bytes using the crypto/rand package
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("%w", err)
	}

	charsetLen := len(charset)

	// If the charset is provided, map the random bytes to the charset
	if charsetLen > 0 {
		for i := 0; i < n; i++ {
			bytes[i] = charset[bytes[i]%byte(charsetLen)]
		}
	}

	return string(bytes), nil
}
