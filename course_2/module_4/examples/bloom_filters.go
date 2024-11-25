package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
)

// BloomFilter represents a Bloom Filter data structure.
type BloomFilter struct {
	bitArray  []byte
	size      int
	numHashes int
}

// NewBloomFilter initializes a new Bloom Filter.
func NewBloomFilter(numItems int, falsePositiveRate float64) (*BloomFilter, error) {
	if numItems <= 0 || falsePositiveRate <= 0 || falsePositiveRate >= 1 {
		return nil, fmt.Errorf("invalid parameters for BloomFilter")
	}

	// Calculate optimal size and number of hashes.
	size := int(math.Ceil(-float64(numItems) * math.Log(falsePositiveRate) / (math.Log(2) * math.Log(2))))
	numHashes := int(math.Max(1, math.Floor((float64(size)/float64(numItems))*math.Log(2))))

	// Initialize bit array.
	return &BloomFilter{
		bitArray:  make([]byte, (size+7)/8), // 8 bits per byte.
		size:      size,
		numHashes: numHashes,
	}, nil
}

// hash calculates a hash value for the given seed.
func (bf *BloomFilter) hash(data string, seed int) int {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	binary.Write(hasher, binary.LittleEndian, uint32(seed))
	hashValue := binary.BigEndian.Uint64(hasher.Sum(nil))
	return int(hashValue % uint64(bf.size))
}

// setBit sets a specific bit in the bit array.
func (bf *BloomFilter) setBit(index int) {
	bf.bitArray[index/8] |= 1 << (index % 8)
}

// getBit checks if a specific bit is set in the bit array.
func (bf *BloomFilter) getBit(index int) bool {
	return bf.bitArray[index/8]&(1<<(index%8)) != 0
}

// Add adds an element to the Bloom Filter.
func (bf *BloomFilter) Add(element string) {
	for i := 0; i < bf.numHashes; i++ {
		index := bf.hash(element, i)
		bf.setBit(index)
	}
}

// Contains checks if an element is possibly in the Bloom Filter.
func (bf *BloomFilter) Contains(element string) bool {
	for i := 0; i < bf.numHashes; i++ {
		index := bf.hash(element, i)
		if !bf.getBit(index) {
			return false
		}
	}
	return true
}

// String provides a string representation of the Bloom Filter.
func (bf *BloomFilter) String() string {
	return fmt.Sprintf("BloomFilter(size=%d, num_hashes=%d)", bf.size, bf.numHashes)
}

// Example usage
func main() {
	// Create a Bloom Filter for 100 items with a 1% false positive rate.
	bloomFilter, err := NewBloomFilter(100, 0.01)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print Bloom Filter info in the required format.
	fmt.Println(bloomFilter)

	// Add elements to the Bloom Filter.
	itemsToAdd := []string{"apple", "banana", "orange", "grape"}
	for _, item := range itemsToAdd {
		bloomFilter.Add(item)
		fmt.Printf("Added: %s\n", item)
	}

	// Check for elements in the Bloom Filter.
	checkItems := []string{"apple", "cherry", "orange", "watermelon"}
	for _, item := range checkItems {
		result := bloomFilter.Contains(item)
		fmt.Printf("Item '%s' is in the Bloom Filter: %t\n", item, result)
	}
}
