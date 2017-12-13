package main

import (
	"encoding/binary"
	"strconv"
	"testing"
	"time"
)

func BenchmarkStrconv(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = []byte(strconv.FormatInt(time.Now().Unix(), 10))
	}
}
func BenchmarkBinary(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, uint64(time.Now().Unix()))
	}
}
