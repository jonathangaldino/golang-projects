package main

import (
	Helpers "colabfiltering/helpers"
	"colabfiltering/similarities"
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")

	personAVector := Helpers.GenerateCombinedVector()
	personBVector := Helpers.GenerateCombinedVector()
	personCVector := Helpers.GenerateCombinedVector()
	personDVector := Helpers.GenerateCombinedVector()

	fmt.Printf("Cosine similarity between person A and person B: %f\n", similarities.CosineSimilarity(personAVector, personBVector))
	fmt.Printf("Cosine similarity between person A and person C: %f\n", similarities.CosineSimilarity(personAVector, personCVector))
	fmt.Printf("Cosine similarity between person A and person D: %f\n", similarities.CosineSimilarity(personAVector, personDVector))
}
