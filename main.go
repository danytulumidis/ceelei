package main

import (
	"ceelei/cmd"
	"fmt"
	"log"

	"github.com/mitchellh/go-homedir"

	bolt "go.etcd.io/bbolt"
)

func main() {
	dir, _ := homedir.Dir()
	dir += "/ceelei.db"
	db, err := bolt.Open(dir, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("ceelei"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	defer db.Close()
	cmd.Execute(db)
}
