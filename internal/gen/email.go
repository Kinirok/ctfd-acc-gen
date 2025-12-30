package gen

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateEmail() string {
	const n = 10
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var email strings.Builder
	email.Grow(n + 12)

	for range n {
		fmt.Fprintf(&email, "%c", charset[seededRand.Intn(10)])
	}
	email.WriteString("@example.com")
	return email.String()
}
