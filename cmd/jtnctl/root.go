package jtnctl

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jtnctl",
	Short: "jtnctl is a CLI tool for simplifying kubectl commands",
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.AddCommand(logsCmd)
	rootCmd.AddCommand(portForwardCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(execCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
