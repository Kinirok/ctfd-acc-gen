package gen

import (
	"fmt"
	"strings"
)

func GenerateEmail() string {
	const n = 10
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var email strings.Builder
	email.Grow(n + 12)

	for range n {
		fmt.Fprintf(&email, "%c", charset[seededRand.Intn(len(charset))])
	}
	email.WriteString("@example.com")
	return email.String()
}
