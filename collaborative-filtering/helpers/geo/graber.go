package Geo

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
)

func OpenGeoJsonFile(filename string) *geojson.FeatureCollection {
	file, err := os.Open(filename)

	// Open the GeoJSON file
	if err != nil {
		fmt.Println("Failed to open the GeoJSON file")
		log.Fatal(err)

	}

	defer file.Close()

	// Read the contents of the file
	byteValue, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("Failed to read the contents of the file")
		log.Fatal(err)
	}

	// Parse the GeoJSON
	fc, _ := geojson.UnmarshalFeatureCollection(byteValue)

	return fc
}

func OpenCityCoordinates(fc *geojson.FeatureCollection) (func() orb.Point, error) {
	features := fc.Features[0]

	mupolygon, ok := features.Geometry.(orb.MultiPolygon)

	if !ok {
		log.Fatal("The geometry is not a multi polygon")
	}

	// Get the bounding box of the mupolygon
	bbox := mupolygon.Bound()

	return func() orb.Point {
		var point orb.Point
		for {
			lng := bbox.Min[0] + rand.Float64()*(bbox.Max[0]-bbox.Min[0])
			lat := bbox.Min[1] + rand.Float64()*(bbox.Max[1]-bbox.Min[1])
			point = orb.Point{lng, lat}

			// Check if the point is inside the polygon
			if planar.MultiPolygonContains(mupolygon, point) {
				break
			}
		}
		return point
	}, nil
}
