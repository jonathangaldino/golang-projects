package collections

import (
	Models "colabfiltering/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DBUser struct {
	Id          string
	Location    bson.D
	Preferences bson.D
}

func MapUserDomainToMongo(user *Models.User) interface{} {
	dbuser := &DBUser{
		Id: user.Id,
		Location: bson.D{
			{Key: "type", Value: "Point"},
			// I don't know if this is going to work with the 2dsphere search index
			{Key: "coordinates", Value: bson.A{
				user.Location.Longitude,
				user.Location.Latitude,
			}},
		},
		Preferences: bson.D{
			{Key: "ethnicities", Value: user.Preferences.Ethnicities},
			{Key: "hobbies", Value: user.Preferences.Hobbies},
			{Key: "favFoods", Value: user.Preferences.FavFoods},
			{Key: "musicGenres", Value: user.Preferences.MusicGenres},
			{Key: "personalityTraits", Value: user.Preferences.PersonalityTraits},
		},
	}
	return dbuser
}

func BulkInsertUsers(client *mongo.Client, users []interface{}) {
	coll := client.Database("maincluster").Collection("users")

	_, err := coll.InsertMany(context.TODO(), users)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result)
}

func InsertUser(client *mongo.Client, user Models.User) {
	coll := client.Database("maincluster").Collection("users")

	result, err := coll.InsertOne(context.TODO(), MapUserDomainToMongo(&user))

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
