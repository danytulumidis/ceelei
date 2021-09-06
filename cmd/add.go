package cmd

import (
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new CLI Command to your list",
	Long:  `Save a new CLI Command into your Database to look it up.
	To look it up use ceelei list`,
	Args: cobra.ExactValidArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
	  addToDatabase(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)  
}

func addToDatabase(args []string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ceelei"))
		err := b.Put([]byte(args[0]), []byte(args[1]))
		return err
	})
}