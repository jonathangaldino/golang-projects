package characteristics

import (
	"math/rand"
	"time"
)

// Given a map of characteristics, generate a random set of characteristics
func GenerateRandomCharacteristics(m map[string]int) []string {
	// Extract keys into a slice
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)

	// Seed the random number generator

	// Number of random keys to select
	n := 3
	if n > len(keys) {
		n = len(keys) // Adjust n if it's greater than the number of available keys
	}

	// Select n random keys
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	randomKeys := keys[:n]

	return randomKeys
}
