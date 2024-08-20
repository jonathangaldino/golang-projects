package main

import (
	"fmt"
	"math"
	"sort"
)

type Person struct {
	Religion int
	Hobby    []int
	Vacation []int
}

type Neighbor struct {
	Index    int
	Distance float64
}

func combineFeatures(person Person, weights map[string]float64) []float64 {
	religion := float64(person.Religion) * weights["religion"]

	hobby := make([]float64, len(person.Hobby))
	for i, h := range person.Hobby {
		hobby[i] = float64(h) * weights["hobby"]
	}

	vacation := make([]float64, len(person.Vacation))
	for i, v := range person.Vacation {
		vacation[i] = float64(v) * weights["vacation"]
	}

	combined := []float64{religion}
	combined = append(combined, hobby...)
	combined = append(combined, vacation...)

	return combined
}

func euclideanDistance(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func knn(train []Person, newPerson Person, k int, weights map[string]float64) []Neighbor {
	combinedNewPerson := combineFeatures(newPerson, weights)

	var neighbors []Neighbor

	for i, person := range train {
		combinedPerson := combineFeatures(person, weights)
		dist := euclideanDistance(combinedPerson, combinedNewPerson)
		neighbors = append(neighbors, Neighbor{Index: i, Distance: dist})
	}

	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].Distance < neighbors[j].Distance
	})

	for i, _ := range neighbors {
		fmt.Printf("Person %d: %f\n", neighbors[i].Index, neighbors[i].Distance)
	}

	return neighbors[:k]
}

func main() {
	// Weights

	weights := map[string]float64{
		"religion": 3,
		"hobby":    1,
		"vacation": 2,
	}

	people := []Person{
		{
			Religion: 4,
			Hobby:    []int{0, 0, 1},
			Vacation: []int{0, 1, 0},
		},
		{
			Religion: 2,
			Hobby:    []int{1, 0, 0},
			Vacation: []int{1, 0, 0},
		},
		{
			Religion: 3,
			Hobby:    []int{0, 0, 1},
			Vacation: []int{0, 1, 0},
		},
	}

	// Using this person to find the K nearest neighbors
	newPerson := Person{Religion: 1, Hobby: []int{0, 1, 0}, Vacation: []int{0, 0, 1}}

	//                 dataset, nn of new person, K, vector weights
	neighbors := knn(people, newPerson, 1, weights)

	fmt.Printf("Index of the closest person: %d\n", neighbors[0].Index)
	fmt.Printf("Distance to the closest person: %f\n", neighbors[0].Distance)
}
