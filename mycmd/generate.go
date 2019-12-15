package main

import (
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use: "generate",
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
