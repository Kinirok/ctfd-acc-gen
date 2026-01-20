package gen

import (
	"fmt"
	"strings"
)

func GenerateTeamName() string {
	const n = 10
	const nums = "0123456789"
	var teamName strings.Builder
	teamName.Grow(n + 5)
	teamName.WriteString("team_")
	for range n {
		fmt.Fprintf(&teamName, "%c", nums[seededRand.Intn(len(nums))])
	}

	return teamName.String()
}
