package utils

import "strings"

func IsDuplicateKeyError(err error) bool {
	// Check if the error string contains a substring indicating a duplicate key error
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}
