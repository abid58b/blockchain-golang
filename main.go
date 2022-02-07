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

	var option uint
	for {
		fmt.Println("Welcome to the Abid Blockchain...")
		fmt.Println("Please Choose Correct Option:")
		fmt.Println("Press 1 for Add Block")
		fmt.Println("Press 2 for View Blockchain:")
		fmt.Println("Press 3 for Exit: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			{
				var data string
				fmt.Println("Please Enter Information / Data You Want to add in Block = ")
				fmt.Scan(&data)
				(chain.AddBlock(data))
				fmt.Println("Block is Successfully Added ")
				continue
			}
		case 2:
			{
				for _, block := range chain.blocks {
					fmt.Printf("Time: %d\n", block.Timestamp)
					fmt.Printf("Previous Hash: %x\n", block.Prevhash)
					fmt.Printf("Data: %s\n", block.Data)
					fmt.Printf("Hash: %x\n", block.Hash)
					fmt.Println("******************")
					fmt.Println("******************")
				}
			}
		case 3:
			break
		default:
			fmt.Println("Wrong Statement")
			
		}

	}

}
