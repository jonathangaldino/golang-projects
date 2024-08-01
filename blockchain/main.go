package main

import (
	Blockchain "blockchain/blockchain"
	"fmt"
	"strconv"
)

func main() {
	chain := Blockchain.InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %x\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)

		pow := Blockchain.NewProof(block)

		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
