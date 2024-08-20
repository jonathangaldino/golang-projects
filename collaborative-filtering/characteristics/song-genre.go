package characteristics

var SongGenreMap = map[string]int{
	"Pop":       0,
	"Reggae":    1,
	"Hip-Hop":   2,
	"Jazz":      3,
	"Rock":      4,
	"Samba":     5,
	"Eletronic": 6,
	"Classic":   7,
	"Blues":     8,
}

var SongGenreMaxFill = 4

func MakeSongGenreVector(userGenres []string) []int {
	songGenreVector := make([]int, len(SongGenreMap))

	for _, genre := range userGenres {
		if index, exists := SongGenreMap[genre]; exists {
			songGenreVector[index] = 1
		}
	}

	return songGenreVector
}
