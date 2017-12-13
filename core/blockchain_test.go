package core

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBlockChain(t *testing.T) {
	assert := assert.New(t)

	chain := NewBlockChain()

	assert.Equal(1, len(chain.Blocks), "Should contain the Genesis Block")
	assert.Equal("Genesis Block", string(chain.Blocks[0].Data))
	assert.Equal([]byte{}, chain.Blocks[0].PrevBlockHash)
}

func TestBlockChainAppendBlock(t *testing.T) {
	assert := assert.New(t)

	hash, _ := hex.DecodeString("d8b3c95bda76107d22eaf64b4ecaa09f496485e1bc7b552e5e28a1004e5fa21f")
	chain := &BlockChain{[]*Block{
		&Block{
			Timestamp:     123456789,
			Data:          []byte("Hello World"),
			PrevBlockHash: []byte(""),
			Hash:          hash,
		},
	}}
	chain.AppendBlock("Ground control to Major Tom")

	assert.Equal(2, len(chain.Blocks))
	assert.Equal("Ground control to Major Tom", string(chain.Blocks[1].Data))
	assert.Equal(hash, chain.Blocks[1].PrevBlockHash)
}
