package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
  }
  
  var versionCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new CLI Command to your list",
	Long:  `Save a new CLI Command into your Database to look it up.
	To look it up use ceelei list`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("Ceelei Add Command coming soon!")
	},
  }