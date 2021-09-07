package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var removeCmd = &cobra.Command{
  Use:   "remove <id>",
  Short: "Remove a command with the specified id",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    id, _ := strconv.ParseInt(args[0], 10, 64)
    removeFromDatabase(int(id))
  },
}

func init() {
  rootCmd.AddCommand(removeCmd)
}

func removeFromDatabase(id int) {
  db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ceelei"))
		
		return b.Delete(itob(id))
	})
}