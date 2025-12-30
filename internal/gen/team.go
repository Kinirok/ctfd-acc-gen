package gen

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateTeamName() string {
	const n = 10
	const nums = "0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var teamName strings.Builder
	teamName.Grow(n + 5)
	teamName.WriteString("team_")
	for range n {
		fmt.Fprintf(&teamName, "%c", nums[seededRand.Intn(10)])
	}

	return teamName.String()
}
