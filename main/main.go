package main

import (
	"exp_blockchain/others"
	"github.com/boltdb/bolt"
)

func main() {
	bc := others.NewBlockchain()
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			panic(nil)
		}
	}(bc.Db)

	cli := others.CLI{bc}
	cli.Run()
}
