package utils

import (
	"crypto/rand"
)

func GenerateRandomString(n int) (string, error) {
	// Define the character set for the random string
	const alphanum string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	bytes := make([]byte, n)
	// Read random bytes from the system's cryptographic random number generator
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		// Convert each byte to a random index within the character set
		bytes[i] = alphanum[int(b)%len(alphanum)]
	}

	return string(bytes), nil
}
