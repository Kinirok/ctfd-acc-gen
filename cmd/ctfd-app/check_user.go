package main

import (
	"log"

	"github.com/spf13/cobra"
)

var checkUserCmd = &cobra.Command{
	Use:   "check_user username1 username2",
	Short: "Проверяет наличие пользователей в базе данных Postgre",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Should be at least 1 username")
			return
		}
		for _, username := range args {
			b, err := generator.CheckTeam(ctx, username)
			if err != nil {
				log.Printf("%s: %s", username, err.Error())

			} else {
				log.Printf("%s: %t", username, b)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkUserCmd)
}
