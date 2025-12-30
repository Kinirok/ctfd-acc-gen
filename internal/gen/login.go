package gen

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateLogin() string {
	const n = 10
	const nums = "0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var login strings.Builder
	login.Grow(n + 5)
	login.WriteString("user_")
	for range n {
		fmt.Fprintf(&login, "%c", nums[seededRand.Intn(10)])
	}

	return login.String()
}
