package gen

import (
	"fmt"
	"strings"
)

func GeneratePassword() string {
	const n = 10
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var password strings.Builder
	password.Grow(n)
	for range n {
		fmt.Fprintf(&password, "%c", charset[seededRand.Intn(len(charset))])
	}

	return password.String()
}
