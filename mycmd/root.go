package mycmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mycmd",
	Short: "a useful tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("welcome to mycmd")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
