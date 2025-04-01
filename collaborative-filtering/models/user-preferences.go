package Models

import "colabfiltering/characteristics"

type Preferences struct {
	Ethnicities       []string
	Hobbies           []string
	FavFoods          []string
	MusicGenres       []string
	PersonalityTraits []string
}

type PreferencesVector struct {
	Ethnicities       []int
	Hobbies           []int
	FavFoods          []int
	MusicGenres       []int
	PersonalityTraits []int
}

func NewUserPreferences(eth []string, hob []string, fav []string, gen []string, per []string) *Preferences {
	return &Preferences{
		Ethnicities:       eth,
		Hobbies:           hob,
		FavFoods:          fav,
		MusicGenres:       gen,
		PersonalityTraits: per,
	}
}

func (c *Preferences) AsVector() *PreferencesVector {
	v := PreferencesVector{
		Ethnicities:       characteristics.MakeEthnicitiesVector(c.Ethnicities),
		Hobbies:           characteristics.MakeHobbiesVector(c.Hobbies),
		FavFoods:          characteristics.MakeFoodVector(c.FavFoods),
		MusicGenres:       characteristics.MakeSongGenreVector(c.MusicGenres),
		PersonalityTraits: characteristics.MakePersonalityTraitsVector(c.PersonalityTraits),
	}

	return &v
}

func (cv *PreferencesVector) AsOneHotVector() []int {
	size := len(cv.Ethnicities) + len(cv.Hobbies) + len(cv.FavFoods) + len(cv.MusicGenres) + len(cv.PersonalityTraits)
	combinedVector := make([]int, size)

	copy(combinedVector, cv.Ethnicities)
	copy(combinedVector[len(cv.Ethnicities):], cv.Hobbies)
	copy(combinedVector[len(cv.Ethnicities)+len(cv.Hobbies):], cv.FavFoods)
	copy(combinedVector[len(cv.Ethnicities)+len(cv.Hobbies)+len(cv.FavFoods):], cv.MusicGenres)
	copy(combinedVector[len(cv.Ethnicities)+len(cv.Hobbies)+len(cv.FavFoods)+len(cv.MusicGenres):], cv.PersonalityTraits)

	return combinedVector

}
