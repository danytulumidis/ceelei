package cmd

import (
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
  Use:   "remove <id>",
  Short: "Remove a command with the specified id",
  Run: func(cmd *cobra.Command, args []string) {
    // removeFromDatabase()
  },
}

func init() {
  rootCmd.AddCommand(removeCmd)
}

// func removeFromDatabase() {
//   db.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("ceelei"))
		
// 		return b.Put(itob(myCommand.ID), encoded)
// 	})
// }