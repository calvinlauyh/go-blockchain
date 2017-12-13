package main

import (
	"fmt"

	"github.com/yuhlau/go-blockchain/core"
)

func main() {
	chain := core.NewBlockChain()

	chain.AppendBlock("Hello World")
	chain.AppendBlock("Ground control to Major Tom")

	for _, blk := range chain.Blocks {
		fmt.Printf("Prev. hash: %x\n", blk.PrevBlockHash)
		fmt.Printf("Data: %s\n", blk.Data)
		fmt.Printf("Hash: %x\n", blk.Hash)
		fmt.Println()
	}
}
