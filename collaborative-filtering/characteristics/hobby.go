package characteristics

var HobbiesMap = map[string]int{
	"Theater":      0,
	"Crafting":     1,
	"Sports":       2,
	"Reading":      3,
	"Gardening":    4,
	"Cooking":      5,
	"Chess":        6,
	"Photography":  7,
	"Dancing":      8,
	"Traveling":    9,
	"Outdoor Show": 10,
}

var HobbiesMaxFill = 3

func MakeHobbiesVector(userHobbies []string) []int {
	vector := make([]int, len(HobbiesMap))

	for _, eth := range userHobbies {
		if index, exists := HobbiesMap[eth]; exists {
			vector[index] = 1
		}
	}

	return vector
}
