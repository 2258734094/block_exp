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

//bc.AddBlock("王显完成了区块链实验的第一步！")
//bc.AddBlock("王显实现了pow")
