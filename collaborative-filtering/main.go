package main

import (
	Config "colabfiltering/Config/env"
	"colabfiltering/database"
	"colabfiltering/database/collections"
	Generators "colabfiltering/helpers/generators"
	Geo "colabfiltering/helpers/geo"
	VectorDB "colabfiltering/vectordb"
)

func main() {
	// Read mongodb URI
	var mongodbUri = Config.MONGODB_URI

	// Connect to Mongodb
	var db = database.NewMongoDB(mongodbUri)
	defer db.Close()

	// Read GeoJSON file
	var fc = Geo.OpenGeoJsonFile("./helpers/geo/cities/rio-de-janeiro.geojson")
	var generator, err = Geo.OpenCityCoordinates(fc)

	if err != nil {
		panic(err)
	}

	vClient, err := VectorDB.NewClient()

	if err != nil {
		panic(err)
	}

	bulkSize := 100
	counter := 0
	list := make([]interface{}, 0, bulkSize)

	for i := 0; i < 1000; i++ {
		coords := generator()
		user := Generators.CreateRandomUser(coords.Lon(), coords.Lat())

		list = append(list, collections.MapUserDomainToMongo(user))
		counter++

		if counter == bulkSize {
			collections.BulkInsertUsers(db.Client, list)
			list = make([]interface{}, 0, bulkSize)
			counter = 0
		}
	}

	if counter > 0 {
		collections.BulkInsertUsers(db.Client, list)
		vClient.BulkSaveEmbeddings(list)
	}

	// coords := generator()
	// user := Generators.CreateRandomUser(coords.Lon(), coords.Lat())

	// collections.InsertUser(db.Client, *user)

	// if err != nil {
	// 	panic(err)
	// }

	// vClient.SaveEmbeddings(*user)

	// queryErr := vClient.QuerySimilarUsers()

	// if queryErr != nil {
	// 	fmt.Println(queryErr)
	// }
}
