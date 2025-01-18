package main

import (
	"fmt"
	"hash"
	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	bitArray []bool 
	hasher hash.Hash32
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		bitArray: make([]bool, size),
		hasher: murmur3.New32(),
	}
}

func (bf *BloomFilter) Add(data []byte) {
	bf.hasher.Write(data)
	hash := bf.hasher.Sum32()
	bf.hasher.Reset()
	fmt.Println(hash)
}

func (bf *BloomFilter) Contains(data []byte) bool{
	bf.hasher.Write(data)
	hash := bf.hasher.Sum32()
	bf.hasher.Reset()
	fmt.Println(hash)
	return false
}

func main() {
	bf := NewBloomFilter(10)
	init_strings := []string{"hello", "world", "foo", "bar", "baz"}
	for _, s := range init_strings {
		bf.Add([]byte(s))
	}
	for _, s := range init_strings {
		bf.Contains([]byte(s))
	}
}