package core

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
}

// NewBlock creates and returns the pointer to a block with the specified data
// and previous block hash using the current time as timestamp
func NewBlock(data string, prevBlockHash []byte) *Block {
	blk := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
	}
	blk.UpdateHash()
	return blk
}

// NewGenesisBlock creates and returns
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// UpdateHash calculates the SHA256 hash based on the Block information and
// updates the Hash field
func (blk *Block) UpdateHash() {
	// Convert timestamp to string representation before converting to byte
	// slice
	timestamp := []byte(strconv.FormatInt(blk.Timestamp, 10))
	headers := bytes.Join([][]byte{
		blk.PrevBlockHash,
		blk.Data,
		timestamp,
	}, []byte{})
	hash := sha256.Sum256(headers)

	blk.Hash = hash[:]
}
