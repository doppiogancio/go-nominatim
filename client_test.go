package go_nominatim

import (
	"fmt"
	"github.com/doppiogancio/go-nominatim/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_JsonV2_Search(t *testing.T) {
	client := NewJsonV2()

	places, err := client.Search(shared.SearchRequest{
		Q:              "Piazza del Plebiscito, Napoli",
		AcceptLanguage: "it",
	})

	assert.Nil(t, err)
	fmt.Println(places[0].DisplayName)
	fmt.Println(places[0].Lat)
	fmt.Println(places[0].Lon)
}

func Test_JsonV2_Reverse(t *testing.T) {
	client := NewJsonV2()

	place, err := client.Reverse(
		"40.8358846",
		"14.2487679",
	)

	assert.Nil(t, err)
	fmt.Println(place.Name)
	fmt.Println(place.DisplayName)
	fmt.Println(place.Address)
}

func Test_GeoJson_Search(t *testing.T) {
	client := NewGeoJson()

	collection, err := client.Search(shared.SearchRequest{
		Q:              "Piazza del Plebiscito, Napoli",
		AcceptLanguage: "it",
	})

	assert.Nil(t, err)
	fmt.Println(collection.Features[0].Properties.DisplayName)
	fmt.Println(collection.Features[0].Geometry.Coordinates)
}

func Test_GeocodeJson_Search(t *testing.T) {
	client := NewGeocodeJson()

	collection, err := client.Search(shared.SearchRequest{
		Q:              "Piazza del Plebiscito, Napoli",
		AcceptLanguage: "it",
	})

	assert.Nil(t, err)
	fmt.Println(collection.Features[0].Properties.Geocoding)
	fmt.Println(collection.Features[0].Properties.Geocoding.Label)
	fmt.Println(collection.Features[0].Geometry.Coordinates)
}
