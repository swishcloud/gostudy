package main

import (
	"fmt"

	"github.com/swishcloud/gostudy/keygenerator"

	"github.com/spf13/cobra"
)

var generateKeyCmd = &cobra.Command{
	Use: "key ",
	Run: func(cmd *cobra.Command, args []string) {
		len, err := cmd.Flags().GetInt32("len")
		Error(err)
		requireUpperCase, err := cmd.Flags().GetBool("ru")
		Error(err)
		requireLowerCase, err := cmd.Flags().GetBool("rl")
		Error(err)
		requireSpecialSymbol, err := cmd.Flags().GetBool("rs")
		Error(err)
		key := keygenerator.NewKey(int(len), requireUpperCase, requireLowerCase, requireSpecialSymbol)
		fmt.Println(key)
	},
}

func init() {
	generateCmd.AddCommand(generateKeyCmd)
	generateKeyCmd.Flags().Int32P("len", "l", 8, "indicate lenth of out key,the default is 8")
	generateKeyCmd.Flags().Bool("ru", false, "indicate whether upper case letter is required,the default is false")
	generateKeyCmd.Flags().Bool("rl", false, "indicate whether lower case letter is required,the default is false")
	generateKeyCmd.Flags().Bool("rs", false, "indicate whether special symbol is required,the default is false")
}
