package characteristics

func CombineFeatures(userEthnicities []int, userHobbies []int, userFavFoods []int, userGenres []int, userPersonalityTraits []int) []int {
	size := len(userEthnicities) + len(userHobbies) + len(userFavFoods) + len(userGenres) + len(userPersonalityTraits)
	combinedVector := make([]int, size)

	copy(combinedVector, userEthnicities)
	copy(combinedVector[len(userEthnicities):], userHobbies)
	copy(combinedVector[len(userEthnicities)+len(userHobbies):], userFavFoods)
	copy(combinedVector[len(userEthnicities)+len(userHobbies)+len(userFavFoods):], userGenres)
	copy(combinedVector[len(userEthnicities)+len(userHobbies)+len(userFavFoods)+len(userGenres):], userPersonalityTraits)

	return combinedVector
}
