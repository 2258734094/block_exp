package main

import "github.com/boltdb/bolt"

func main() {
	bc := NewBlockchain()
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {

		}
	}(bc.db)

	cli := CLI{bc}
	cli.Run()
}
