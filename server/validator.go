package main

import "regexp"

func sixToEightDigitAlphanumericPasswordValidator(password string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]{6,8}$`)
	return re.MatchString(password)
}
