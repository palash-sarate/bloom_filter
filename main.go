package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
	"hash"
)

// TODO: add more hash functions to reduce false positives

type BloomFilter struct {
	size       int
	bitArray   []bool
	numHashers int
	hasher     hash.Hash32
	hashers    []hash.Hash32
}

func NewBloomFilter(size int, numHashers int, hashers []hash.Hash32) *BloomFilter {
	if numHashers == 0 && len(hashers) == 0 {
		numHashers = 1
		for i := 0; i < numHashers; i++ {
			hashers = append(hashers, murmur3.New32WithSeed(uint32(i)))
		}
	}

	if numHashers != 0 && len(hashers) == 0 {
		for i := 0; i < numHashers; i++ {
			hashers = append(hashers, murmur3.New32WithSeed(uint32(i)))
		}
	}

	if numHashers == 0 && len(hashers) != 0 {
		numHashers = len(hashers)
	}

	return &BloomFilter{
		size:       size,
		bitArray:   make([]bool, size),
		hasher:     murmur3.New32(),
		numHashers: numHashers,
		hashers:    hashers,
	}
}

func (bf *BloomFilter) Add(data []byte) {
	for _, hasher := range bf.hashers {
		// hash the data
		hasher.Write(data)
		hash := hasher.Sum32()
		// get the index of the bit array from the hash
		index := int(hash) % bf.size
		bf.bitArray[index] = true
		// reset the hasher
		hasher.Reset()
	}
}

func (bf *BloomFilter) Contains(data []byte) bool {
	for _, hasher := range bf.hashers {
		// hash the data
		hasher.Write(data)
		hash := hasher.Sum32()
		// get the index of the bit array from the hash
		index := int(hash) % bf.size
		hasher.Reset()
		if !bf.bitArray[index] {
            return false
        }
	}
	return true
}

func main() {
	bf := NewBloomFilter(100, 3, nil)
	init_strings := []string{"hello", "world", "foo", "bar", "baz"}
	check_strings := []string{"hello", "world", "foo", "bar", "palash"}

	for _, s := range init_strings {
		bf.Add([]byte(s))
	}
	for _, s := range check_strings {
		fmt.Println(bf.Contains([]byte(s)))
	}
}
