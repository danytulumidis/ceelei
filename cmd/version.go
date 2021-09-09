package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the current version number of Ceelei",
  Long:  `The version of Ceelei that is currently installed on your system`,
  Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("Ceelei v0.1")
  },
}

func init() {
  rootCmd.AddCommand(versionCmd)
}
