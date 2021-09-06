package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
  Use:   "list",
  Short: "List all your saved CLI commands",
  Long:  `A list of all your saved CLI commands that you saved within Ceelei currently`,
  Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("List Command coming soon!")
  },
}

func init() {
  rootCmd.AddCommand(listCmd)
}
