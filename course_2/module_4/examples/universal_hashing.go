package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// KeyValue represents a single key-value pair in the hash table.
type KeyValue struct {
	Key   []int
	Value interface{}
}

// UniversalHashFunction represents a universal hash function.
type UniversalHashFunction struct {
	p      int        // A prime number larger than the universe size.
	m      int        // Number of buckets in the hash table.
	a      []int      // Random coefficients for the hash function.
	b      int        // Random constant for the hash function.
	k      int        // Dimensionality of the keys.
	random *rand.Rand // Random number generator for reproducibility.
}

// NewUniversalHashFunction initializes a new universal hash function.
func NewUniversalHashFunction(p, m, k int) (*UniversalHashFunction, error) {
	if !isPrime(p) {
		return nil, errors.New("p must be a prime number")
	}

	// Use a local random number generator for thread safety and reproducibility.
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	// Generate random coefficients a and b.
	a := make([]int, k)
	for i := 0; i < k; i++ {
		a[i] = rnd.Intn(p-1) + 1
	}
	b := rnd.Intn(p)

	return &UniversalHashFunction{
		p:      p,
		m:      m,
		a:      a,
		b:      b,
		k:      k,
		random: rnd,
	}, nil
}

// Hash computes the hash value for a given key.
func (uhf *UniversalHashFunction) Hash(key []int) (int, error) {
	// Ensure the key has the correct dimensionality.
	if len(key) != uhf.k {
		return -1, fmt.Errorf("key must have size %d, got %d", uhf.k, len(key))
	}

	// Ensure all key values are within the valid range [0, p).
	for _, x := range key {
		if x < 0 || x >= uhf.p {
			return -1, fmt.Errorf("key value %d out of range [0, %d)", x, uhf.p)
		}
	}

	// Compute the hash value using the formula: ((Σ(a[i] * key[i]) + b) % p) % m.
	sum := uhf.b
	for i := 0; i < len(key); i++ {
		sum += uhf.a[i] * key[i]
	}
	return (sum % uhf.p) % uhf.m, nil
}

// HashTableWithChaining represents a hash table with chaining for collision resolution.
type HashTableWithChaining struct {
	numBuckets          int                    // Number of buckets in the hash table.
	loadFactorThreshold float64                // Load factor threshold for resizing.
	size                int                    // Current number of elements in the table.
	table               [][]KeyValue           // Buckets represented as slices of KeyValue pairs.
	hashFunction        *UniversalHashFunction // Universal hash function used for hashing.
}

// NewHashTableWithChaining initializes a new hash table with chaining.
func NewHashTableWithChaining(numBuckets, p, k int, loadFactorThreshold float64) (*HashTableWithChaining, error) {
	hashFunc, err := NewUniversalHashFunction(p, numBuckets, k)
	if err != nil {
		return nil, err
	}

	// Initialize the hash table with empty buckets.
	table := make([][]KeyValue, numBuckets)
	return &HashTableWithChaining{
		numBuckets:          numBuckets,
		loadFactorThreshold: loadFactorThreshold,
		size:                0,
		table:               table,
		hashFunction:        hashFunc,
	}, nil
}

// Insert adds a key-value pair to the hash table.
func (ht *HashTableWithChaining) Insert(key []int, value interface{}) error {
	// Ensure the key has the correct dimensionality.
	if len(key) != ht.hashFunction.k {
		return fmt.Errorf("key must have size %d", ht.hashFunction.k)
	}

	// Resize the table if the load factor exceeds the threshold.
	if float64(ht.size)/float64(ht.numBuckets) > ht.loadFactorThreshold {
		if err := ht.resize(); err != nil {
			return err
		}
	}

	// Compute the bucket index for the key.
	bucketIndex, err := ht.hashFunction.Hash(key)
	if err != nil {
		return err
	}

	// Check if the key already exists and update its value if so.
	for i, pair := range ht.table[bucketIndex] {
		if equalKeys(pair.Key, key) {
			ht.table[bucketIndex][i].Value = value
			return nil
		}
	}

	// Add the new key-value pair to the bucket.
	ht.table[bucketIndex] = append(ht.table[bucketIndex], KeyValue{Key: key, Value: value})
	ht.size++
	return nil
}

// Search retrieves a value by its key.
func (ht *HashTableWithChaining) Search(key []int) (interface{}, error) {
	// Compute the bucket index for the key.
	bucketIndex, err := ht.hashFunction.Hash(key)
	if err != nil {
		return nil, err
	}

	// Search for the key in the bucket.
	for _, pair := range ht.table[bucketIndex] {
		if equalKeys(pair.Key, key) {
			return pair.Value, nil
		}
	}
	return nil, nil // Key not found.
}

// Delete removes a key-value pair from the hash table.
func (ht *HashTableWithChaining) Delete(key []int) (bool, error) {
	// Compute the bucket index for the key.
	bucketIndex, err := ht.hashFunction.Hash(key)
	if err != nil {
		return false, err
	}

	// Search for the key in the bucket and remove it.
	for i, pair := range ht.table[bucketIndex] {
		if equalKeys(pair.Key, key) {
			ht.table[bucketIndex] = append(ht.table[bucketIndex][:i], ht.table[bucketIndex][i+1:]...)
			ht.size--
			return true, nil
		}
	}
	return false, nil // Key not found.
}

// resize doubles the size of the hash table and rehashes all elements.
func (ht *HashTableWithChaining) resize() error {
	oldTable := ht.table
	newBuckets := ht.numBuckets * 2

	// Create a new hash function with updated bucket count.
	hashFunc, err := NewUniversalHashFunction(ht.hashFunction.p, newBuckets, ht.hashFunction.k)
	if err != nil {
		return err
	}

	// Reinitialize the table and reinsert all elements.
	ht.numBuckets = newBuckets
	ht.table = make([][]KeyValue, newBuckets)
	ht.hashFunction = hashFunc
	ht.size = 0

	for _, bucket := range oldTable {
		for _, pair := range bucket {
			if err := ht.Insert(pair.Key, pair.Value); err != nil {
				return err
			}
		}
	}
	return nil
}

// Display prints the contents of the hash table.
func (ht *HashTableWithChaining) Display(limit int) {
	fmt.Printf("Hash Table (size: %d, buckets: %d):\n", ht.size, ht.numBuckets)
	displayed := 0
	for i, bucket := range ht.table {
		if len(bucket) > 0 {
			fmt.Printf("Bucket %d: %v\n", i, bucket)
			displayed++
		}
		if displayed >= limit {
			fmt.Println("... and more buckets.")
			break
		}
	}
}

// Helper Functions

// isPrime checks if a number is prime.
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// equalKeys checks if two keys are equal.
func equalKeys(key1, key2 []int) bool {
	if len(key1) != len(key2) {
		return false
	}
	for i := 0; i < len(key1); i++ {
		if key1[i] != key2[i] {
			return false
		}
	}
	return true
}

// Main function for testing the hash table.
func main() {
	// Use a prime number larger than the maximum value in keys (e.g., 257 for IP addresses)
	hashTable, err := NewHashTableWithChaining(10, 257, 4, 0.7)
	if err != nil {
		fmt.Println("Error creating hash table:", err)
		return
	}

	// Insert key-value pairs with error handling
	if err := hashTable.Insert([]int{192, 168, 0, 1}, "Device A"); err != nil {
		fmt.Println("Error inserting key [192, 168, 0, 1]:", err)
	}
	if err := hashTable.Insert([]int{192, 168, 0, 2}, "Device B"); err != nil {
		fmt.Println("Error inserting key [192, 168, 0, 2]:", err)
	}
	if err := hashTable.Insert([]int{10, 0, 0, 1}, "Device C"); err != nil {
		fmt.Println("Error inserting key [10, 0, 0, 1]:", err)
	}

	// Search for values
	value, err := hashTable.Search([]int{192, 168, 0, 1})
	if err != nil {
		fmt.Println("Error searching key [192, 168, 0, 1]:", err)
	} else {
		fmt.Println("Search [192, 168, 0, 1]:", value)
	}

	value, err = hashTable.Search([]int{192, 168, 0, 3})
	if err != nil {
		fmt.Println("Error searching key [192, 168, 0, 3]:", err)
	} else {
		fmt.Println("Search [192, 168, 0, 3]:", value)
	}

	// Delete a value
	deleted, err := hashTable.Delete([]int{192, 168, 0, 1})
	if err != nil {
		fmt.Println("Error deleting key [192, 168, 0, 1]:", err)
	} else {
		fmt.Println("Delete [192, 168, 0, 1]:", deleted)
	}

	// Display the hash table
	hashTable.Display(5)
}
