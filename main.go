package main

import (
	"fmt"
)

type BloomFilter struct {
	bitArray []bool
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		bitArray: make([]bool, size),
	}
}

func main() {
	bf := NewBloomFilter(10)
	fmt.Println(bf)
}