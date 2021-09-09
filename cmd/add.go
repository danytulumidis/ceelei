package cmd

import (
	"encoding/binary"
	"encoding/json"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

type NewCommand struct {
	ID int
	Command string
	Description string
}

var addCmd = &cobra.Command{
	Use:   "add <command> <description>",
	Short: "Add a new CLI Command to your list.",
	Long:  `Save a new CLI Command into your Database to look it up.
To look it up use ceelei list.`,
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

		id, _ := b.NextSequence()
		myCommand := NewCommand{
			int(id),
			string(args[0]),
			args[1],
		}

		encoded, err := json.Marshal(myCommand)

		if err != nil {
            return err
        }

		return b.Put(itob(myCommand.ID), encoded)
	})
}

func itob(v int) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(v))
    return b
}