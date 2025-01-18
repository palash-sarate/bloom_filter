package bloomFilter

import (
	"github.com/spaolacci/murmur3"
	"hash"
)

type BloomFilter struct {
	size       int
	bitArray   []bool
	numHashers int
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
		println("Correcting numHashers to ", numHashers)
	}

	return &BloomFilter{
		size:       size,
		bitArray:   make([]bool, size),
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


