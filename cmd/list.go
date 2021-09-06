package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var listCmd = &cobra.Command{
  Use:   "list",
  Short: "List all your saved CLI commands",
  Long:  `A list of all your saved CLI commands that you saved within Ceelei currently`,
  Run: func(cmd *cobra.Command, args []string) {
    listFromDatabase()
  },
}

func init() {
  rootCmd.AddCommand(listCmd)
}

func listFromDatabase() {
  db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("ceelei"))
    b.ForEach(func(k, v []byte) error {
      fmt.Printf("Command: %s, Description: %s\n", k, v)
      return nil
    })
    return nil
  })
}