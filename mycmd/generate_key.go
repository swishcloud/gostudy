package main

import (
	"fmt"
	"strconv"

	"github.com/swishcloud/gostudy/keygenerator"

	"github.com/spf13/cobra"
)

var generateKeyCmd = &cobra.Command{
	Use:   "key ",
	Short: "randomly generate string key",
	Long:  "randomly generate string key,the length of out key defaults to 8",
	Run: func(cmd *cobra.Command, args []string) {
		len, err := strconv.Atoi(args[0])
		if err != nil {
			len = 8
		}
		excludeDigits, err := cmd.Flags().GetBool("ed")
		Error(err)
		excludeUpperCase, err := cmd.Flags().GetBool("eu")
		Error(err)
		excludeLowerCase, err := cmd.Flags().GetBool("el")
		Error(err)
		excludeSpecialSymbol, err := cmd.Flags().GetBool("es")
		Error(err)
		key, err := keygenerator.NewKey(int(len), excludeDigits, excludeUpperCase, excludeLowerCase, excludeSpecialSymbol)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(key)
		}
	},
}

func init() {
	generateCmd.AddCommand(generateKeyCmd)
	generateKeyCmd.Flags().Bool("ed", false, "indicate whether exclude upper case letter,the default is false")
	generateKeyCmd.Flags().Bool("eu", false, "indicate whether exclude upper case letter,the default is false")
	generateKeyCmd.Flags().Bool("el", false, "indicate whether exclude lower case letter,the default is false")
	generateKeyCmd.Flags().Bool("es", true, "indicate whether exclude special symbol,the default is true")
}
