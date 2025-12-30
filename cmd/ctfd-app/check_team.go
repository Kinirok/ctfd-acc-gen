package main

import (
	"log"

	"github.com/spf13/cobra"
)

var checkTeamCmd = &cobra.Command{
	Use:   "check_user",
	Short: "Проверяет наличие команд в базе данных Postgre",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Should be at least 1 team name")
			return
		}
		for _, team_name := range args {
			b, err := generator.CheckTeam(ctx, team_name)
			if err != nil {
				log.Printf("%s: %s", team_name, err.Error())

			} else {
				log.Printf("%s: %t", team_name, b)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkTeamCmd)
}
