// Cosine Similarity
// https://en.wikipedia.org/wiki/Cosine_similarity

package similarities

import (
	"math"
)

func CosineSimilarity(vecA, vecB []int) float64 {
	if len(vecA) != len(vecB) {
		panic("Vectors must be of the same size")
	}

	var dotProduct, magnitudeA, magnitudeB float64

	for i := 0; i < len(vecA); i++ {
		// fmt.Printf("vecA[%d]: %d | vecB[%d]: %d\n", i, vecA[i], i, vecB[i])
		dotProduct += float64(vecA[i] * vecB[i])
		magnitudeA += float64(vecA[i] * vecA[i])
		magnitudeB += float64(vecB[i] * vecB[i])
	}

	// fmt.Printf("dotProduct: %f\n", dotProduct)
	// fmt.Printf("magnitudeA: %f\n", magnitudeA)
	// fmt.Printf("magnitudeB: %f\n", magnitudeB)

	// Calculate the magnitudes
	magnitudeA = math.Sqrt(magnitudeA)
	magnitudeB = math.Sqrt(magnitudeB)

	// Handle the case where magnites are zero (avoid division by zero)
	// Also handle the case where dot product is zero (no overlap)
	// Cuz 0/anything is 0
	if magnitudeA == 0 || magnitudeB == 0 || dotProduct == 0 {
		return 0
	}

	return dotProduct / magnitudeA * magnitudeB
}
