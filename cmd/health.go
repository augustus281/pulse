package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/augustus281/pulse/internal/health"
)

var url string

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Check the health of a service",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			fmt.Println("! Please provide a URL to check the health of a service.")
			return
		}
		if err := health.CheckHealth(url); err != nil {
			fmt.Printf("! Health check failed: %v\n", err)
			return
		}
	},
}

func init() {
	healthCmd.Flags().StringVarP(&url, "url", "u", "", "URL to check the health of a service")
	rootCmd.AddCommand(healthCmd)
}
