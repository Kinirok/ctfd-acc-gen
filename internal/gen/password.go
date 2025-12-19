package gen

import (
	"math/rand"
	"time"
)

func generatePassword() string {
	n := 10
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	password := make([]byte, n)

	for i := range password {
		password[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(password)
}
