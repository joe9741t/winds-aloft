package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "winds-aloft",
	Short: "Winds Aloft is a CLI tool",
	Long:  `A longer description of Winds Aloft CLI tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Winds Aloft CLI tool")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show winds aloft for an airport",
	Long:  `Show detailed winds aloft for a specific airport.`,
	Run: func(cmd *cobra.Command, args []string) {
		airport, _ := cmd.Flags().GetString("airport")
		fmt.Printf("Showing information for airport: %s\n", airport)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("airport", "a", "", "Airport code")
	showCmd.MarkFlagRequired("airport")
}

func main() {
	Execute()
}
