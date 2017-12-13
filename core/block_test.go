package core

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewBlock(t *testing.T) {
	assert := assert.New(t)

	be4Time := time.Now().Unix()
	blk := NewBlock("Hello World", []byte{})
	assert.Equal(true, blk.Timestamp >= be4Time)
	assert.Equal("Hello World", string(blk.Data))
	assert.NotEqual([]byte{}, blk.Hash)
}

func TestNewGenesisBlock(t *testing.T) {
	assert := assert.New(t)

	be4Time := time.Now().Unix()
	blk := NewGenesisBlock()
	assert.Equal(true, blk.Timestamp >= be4Time)
	assert.Equal("Genesis Block", string(blk.Data))
	assert.NotEqual([]byte{}, blk.Hash)
}

func TestBlockUpdateHash(t *testing.T) {
	assert := assert.New(t)

	blk := &Block{
		Timestamp:     123456789,
		Data:          []byte("Hello World"),
		PrevBlockHash: []byte(""),
		Hash:          []byte{},
	}

	blk.UpdateHash()
	assert.Equal("d8b3c95bda76107d22eaf64b4ecaa09f496485e1bc7b552e5e28a1004e5fa21f", hex.EncodeToString(blk.Hash))
}
