package characteristics

var EthnicitiesMap = map[string]int{
	"White":    0,
	"Native":   1,
	"Asian":    2,
	"Afro":     3,
	"South":    4,
	"Arab":     5,
	"Indo":     6,
	"Jewish":   7,
	"Tibetan":  8,
	"Brown":    9,
	"Bengalí":  10,
	"Hispanic": 11,
	"Bantu":    12,
}

var EthnicitiesMaxFill = 2

func MakeEthnicitiesVector(userEthnicities []string) []int {
	ethVector := make([]int, len(SongGenreMap))

	for _, eth := range userEthnicities {
		if index, exists := EthnicitiesMap[eth]; exists {
			ethVector[index] = 1
		}
	}

	return ethVector
}
