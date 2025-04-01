package Generators

import (
	"colabfiltering/characteristics"
	Models "colabfiltering/models"

	"github.com/google/uuid"
)

func CreateRandomUser(long, lat float64) *Models.User {
	var user = &Models.User{
		Id:       uuid.New().String(),
		Location: Models.Coordinates{Longitude: long, Latitude: lat},
	}

	eth := characteristics.GenerateRandomCharacteristics(characteristics.EthnicitiesMap)
	hobbies := characteristics.GenerateRandomCharacteristics(characteristics.HobbiesMap)
	favFoods := characteristics.GenerateRandomCharacteristics(characteristics.FavFoodsMap)
	genres := characteristics.GenerateRandomCharacteristics(characteristics.SongGenreMap)
	personalityTraits := characteristics.GenerateRandomCharacteristics(characteristics.PersonalityTraitsMap)

	userPreferences := Models.NewUserPreferences(eth, hobbies, favFoods, genres, personalityTraits)
	user.Preferences = *userPreferences

	return user
}
