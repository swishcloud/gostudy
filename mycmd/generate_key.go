package mycmd

import (
	"fmt"
	"gostudy/keygenerator"

	"github.com/spf13/cobra"
)

var generateKeyCmd = &cobra.Command{
	Use:   "key ",
	Short: "randomly generate string key",
	Long:  "randomly generate string key,the length of out key defaults to 8",
	Run: func(cmd *cobra.Command, args []string) {
		len, err := cmd.Flags().GetInt("len")
		Error(err)
		excludeDigits, err := cmd.Flags().GetBool("ed")
		Error(err)
		excludeUpperCase, err := cmd.Flags().GetBool("eu")
		Error(err)
		excludeLowerCase, err := cmd.Flags().GetBool("el")
		Error(err)
		excludeSpecialSymbol, err := cmd.Flags().GetBool("es")
		Error(err)
		key, err := keygenerator.NewKey(len, excludeDigits, excludeUpperCase, excludeLowerCase, excludeSpecialSymbol)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(key)
		}
	},
}

func init() {
	generateCmd.AddCommand(generateKeyCmd)
	generateKeyCmd.Flags().IntP("len", "l", 8, "indicate the lenth of out key,the default is 8")
	generateKeyCmd.Flags().Bool("ed", false, "indicate whether exclude upper case letter,the default is false")
	generateKeyCmd.Flags().Bool("eu", false, "indicate whether exclude upper case letter,the default is false")
	generateKeyCmd.Flags().Bool("el", false, "indicate whether exclude lower case letter,the default is false")
	generateKeyCmd.Flags().Bool("es", false, "indicate whether exclude special symbol,the default is true")
}
