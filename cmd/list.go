package cmd

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var listCmd = &cobra.Command{
  Use:   "list",
  Short: "List all your saved CLI commands.",
  Long:  `A list of all your saved CLI commands that you saved within Ceelei currently.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("----------CEELEI LIST----------")
    listFromDatabase()
    fmt.Println("-------------------------------")
  },
}

func init() {
  rootCmd.AddCommand(listCmd)
}

func listFromDatabase() {
  db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("ceelei"))

    b.ForEach(func(k, v []byte) error {
      var currCommand NewCommand
      json.Unmarshal(v, &currCommand)
      fmt.Printf("%v: Command: %v - Description: %s\n", btoi(k), currCommand.Command, currCommand.Description)
      return nil
    })
    
    return nil
  })
}

func btoi(b []byte) int {
  return int(binary.BigEndian.Uint64(b))
}