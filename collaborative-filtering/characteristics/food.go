package characteristics

var FavFoodsMap = map[string]int{
	"Japanese":   0,
	"Libanese":   1,
	"Italian":    2,
	"Tailandese": 3,
	"Spanish":    4,
	"Chinese":    5,
	"Mexican":    6,
	"French":     7,
	"Indian":     8,
	"Brazilian":  9,
}

var FoodMaxFill = 2

func MakeFoodVector(userFavFoods []string) []int {
	vector := make([]int, len(FavFoodsMap))

	for _, choice := range userFavFoods {
		if index, exists := FavFoodsMap[choice]; exists {
			vector[index] = 1
		}
	}

	return vector
}
