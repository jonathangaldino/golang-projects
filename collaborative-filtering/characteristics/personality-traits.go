package characteristics

var PersonalityTraitsMap = map[string]int{
	"Pacifist":             0,
	"Good listener":        1,
	"Devout":               2,
	"Dreamer":              3,
	"Reformer":             4,
	"Spiritual Adventurer": 5,
	"Empathetic":           6,
	"Creative":             7,
	"Traditionalist":       8,
	"Spiritual Leader":     9,
	"Athlete":              10,
	"Traveler":             11,
	"Witty":                12,
	"Companion":            13,
	"Entrepeneur":          14,
	"Artistic":             15,
	"Progressive":          16,
	"Funny":                17,
	"Conservative":         18,
	"Intellectual":         19,
	"Spontaneous":          20,
	"Communicator":         21,
}

var PersonalityTraitsMaxFill = 5

func MakePersonalityTraitsVector(userPersonalityTraits []string) []int {
	vector := make([]int, len(PersonalityTraitsMap))

	for _, choice := range userPersonalityTraits {
		if index, exists := PersonalityTraitsMap[choice]; exists {
			vector[index] = 1
		}
	}

	return vector
}
