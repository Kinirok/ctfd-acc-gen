package main

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var createUsersCmd = &cobra.Command{
	Use:   "create_users",
	Short: "Создать пользователей",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("there should be 1 integer argument")
		}
		n, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("argument n must be an integer:", err)
		}
		err = generator.CreateIndividualAccounts(ctx, generator.GenerateNEmails(n), false, nil, nil)
		if err != nil {
			log.Print(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(createUsersCmd)
}
