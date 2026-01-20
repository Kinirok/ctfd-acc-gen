package gen

import (
	"fmt"
	"strings"
)

func GenerateLogin() string {
	const n = 10
	const nums = "0123456789"

	var login strings.Builder
	login.Grow(n + 5)
	login.WriteString("user_")
	for range n {
		fmt.Fprintf(&login, "%c", nums[seededRand.Intn(len(nums))])
	}

	return login.String()
}
