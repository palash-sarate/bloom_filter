package main

import (
	"fmt"
	"hash"
	"github.com/spaolacci/murmur3"
)

// TODO: add more hash functions to reduce false positives

type BloomFilter struct {
	size int
	bitArray []bool 
	hasher hash.Hash32
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		size: size,
		bitArray: make([]bool, size),
		hasher: murmur3.New32(),
	}
}

func (bf *BloomFilter) Add(data []byte) {
	// hash the data
	bf.hasher.Write(data)
	hash := bf.hasher.Sum32()
	// get the index of the bit array from the hash
	index := int(hash) % bf.size
	bf.bitArray[index] = true
	// reset the hasher
	bf.hasher.Reset()
}

func (bf *BloomFilter) Contains(data []byte) bool{
	// hash the data
	bf.hasher.Write(data)
	hash := bf.hasher.Sum32()
	// get the index of the bit array from the hash
	index := int(hash) % bf.size
	bf.hasher.Reset()
	return bf.bitArray[index]
}

func main() {
	bf := NewBloomFilter(100)
	init_strings := []string{"hello", "world", "foo", "bar", "baz"}
	check_strings := []string{"hello", "world", "foo", "bar", "palash"}

	for _, s := range init_strings {
		bf.Add([]byte(s))
	}
	for _, s := range check_strings {
		fmt.Println(bf.Contains([]byte(s)))
	}
}