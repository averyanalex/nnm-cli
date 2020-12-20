package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "nnm-cli",
	Short: "NNM-CLI is a CLI client for NNM",
	Long: `There is no description, sorry`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, user")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
