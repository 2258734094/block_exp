package others

import (
	"github.com/boltdb/bolt"
)

//type Blockchain struct {
//	blocks []*others
//}

type Blockchain struct {
	tip []byte
	Db  *bolt.DB
}

//func (bc *Blockchain) AddBlock(data string) {
//	prevBlock := bc.blocks[len(bc.blocks)-1]
//	newBlock := NewBlock(data, prevBlock.Hash)
//	bc.blocks = append(bc.blocks, newBlock)
//}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocksBucket"))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocksBucket"))
		err := b.Put(newBlock.Hash, newBlock.Serialize())

		if err != nil {
			panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//func NewBlockchain() *Blockchain {
//	return &Blockchain{[]*others{NewGenesisBlock()}} //创建一个新的指针切片，并向其中添加一个元素。这个元素是由 NewGenesisBlock() 函数返回的新创世区块的地址。
//}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open("dbFile", 0600, nil)

	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocksBucket"))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte("blocksBucket"))

			if err != nil {
				panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}
