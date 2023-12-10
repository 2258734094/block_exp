package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))                       //这里使用了 strconv.FormatInt 函数将整数类型的时间戳转换为十进制格式的字符串，然后将该字符串转换为字节切片。
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) //这里使用了 bytes.Join 函数将多个字节切片连接在一起，形成一个新的字节切片。
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

//func NewBlock(data string, prevBlockHash []byte) *others {
//	block := &others{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
//
//	block.SetHash()
//	//SetHash 方法的接收者是 *others 类型，即一个指向 others 结构体的指针。这意味着当你调用 block.SetHash() 时，Go 语言会将 block 指针作为参数传递给 SetHash 方法。在方法内部，你可以直接通过 b 变量（即接收者）来访问和修改 block 指向的 others 结构体实例的所有字段。
//	//
//	//因此，即使没有显式地传入参数，SetHash 方法仍然能够获取到当前 others 结构体实例，并为其计算和设置哈希值。这是因为 SetHash 方法实际上已经接收到一个隐含的参数——指向当前 others 实例的指针。
//	return block
//}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
