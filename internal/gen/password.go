package gen

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GeneratePassword() string {
	const n = 10
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var password strings.Builder
	password.Grow(n)
	for range n {
		fmt.Fprintf(&password, "%c", charset[seededRand.Intn(len(charset))])
	}

	return password.String()
}
