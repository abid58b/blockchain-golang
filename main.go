package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte
	Hash      []byte
	Prevhash  []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	info := bytes.Join([][]byte{timestamp, b.Data, b.Prevhash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), []byte{}, previousHash}
	block.SetHash()
	return block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)

}

func GenesisBlock() *Block {
	return CreateBlock("Genesis", []byte{})

}

func AbidBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func main() {

	chain := AbidBlockchain()

	chain.AddBlock("BTC")
	chain.AddBlock("Solana")
	chain.AddBlock("Eth")

	for _, block := range chain.blocks {
		fmt.Printf("Time: %d\n", block.Timestamp)
		fmt.Printf("Previous Hash: %x\n", block.Prevhash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("******************")
		fmt.Println("******************")
	}

}
