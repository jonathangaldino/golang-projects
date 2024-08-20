package Helpers

import (
	characteristics "colabfiltering/characteristics"
	"math/rand"
)

func GenerateRandomVector(size int, maxFill int) []int {
	if maxFill > size {
		panic("maxFill must be less than or equal to size")
	}

	vector := make([]int, size)

	positions := rand.Perm(size)

	for i := 0; i < maxFill; i++ {
		vector[positions[i]] = 1
	}

	return vector
}

func GenerateCombinedVector() []int {
	totalSize := len(characteristics.EthnicitiesMap) + len(characteristics.HobbiesMap) + len(characteristics.FavFoodsMap) + len(characteristics.SongGenreMap) + len(characteristics.PersonalityTraitsMap)
	maxFill := characteristics.EthnicitiesMaxFill + characteristics.HobbiesMaxFill + characteristics.FoodMaxFill + characteristics.SongGenreMaxFill + characteristics.PersonalityTraitsMaxFill

	vector := GenerateRandomVector(totalSize, maxFill)

	return vector
}
