package main

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var createTeamsCmd = &cobra.Command{
	Use:   "create_teams",
	Short: "Создать команды",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("Invalid arguments, should be 2 integers")
			return
		}
		n, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("argument must be an integer:", err)
			return
		}
		m, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("argument must be an integer:", err)
			return
		}
		err = generator.CreateTeamAccounts(ctx, n, m)
		if err != nil {
			log.Print(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(createTeamsCmd)
}
