package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "pulse",
	Short: "Pulse is a health check tool",
}

func Execute() error {
	return rootCmd.Execute()
}
