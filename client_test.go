package go_nominatim

import (
	"github.com/doppiogancio/go-nominatim/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeocode(t *testing.T) {
	coordinate, err := Geocode("Capri")
	assert.Nil(t, err)
	assert.Equal(t, 40.54877475, coordinate.Latitude)
	assert.Equal(t, 14.22808744705355, coordinate.Longitude)
}

func TestReverseGeocode(t *testing.T) {
	address, err := ReverseGeocode(
		52.5162024,
		13.3777343363579,
		"it",
	)

	assert.Nil(t, err)

	expected := &shared.Address{
		DisplayName: "Porta di Brandeburgo, 1, Pariser Platz, Mitte, Berlino, 10117, Germania",
		Road:        "Pariser Platz",
		Suburb:      "Mitte",
		City:        "Berlino",
		County:      "",
		State:       "Berlino",
		Postcode:    "10117",
		Country:     "Germania",
		CountryCode: "de",
	}

	assert.Equal(t, expected, address)
}

func Test_JsonV2_Search(t *testing.T) {
	places, err := NewJsonV2().Search(shared.SearchRequest{
		Q:              "Piazza del Plebiscito, Napoli",
		AcceptLanguage: "it",
	})

	assert.Nil(t, err)

	assert.Len(t, places, 2)

	assert.Equal(t, "Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia", places[0].DisplayName)
	assert.Equal(t, "40.835855949999996", places[0].Lat)
	assert.Equal(t, "14.248565182098474", places[0].Lon)
}

func Test_JsonV2_Reverse(t *testing.T) {
	_, err := NewJsonV2().Reverse(
		shared.ReverseGeocodeRequest{
			Latitude:       40.8358846,
			Longitude:      14.2487679,
			AcceptLanguage: "it",
		},
	)

	assert.Nil(t, err)
	// TODO assert result
}

func Test_GeoJson_Search(t *testing.T) {
	_, err := NewGeoJson().Search(shared.SearchRequest{
		Q:              "Piazza del Plebiscito, Napoli",
		AcceptLanguage: "it",
	})

	assert.Nil(t, err)
	// TODO assert result
}

func Test_GeocodeJson_Search(t *testing.T) {
	_, err := NewGeocodeJson().Search(shared.SearchRequest{
		Q:              "Piazza del Plebiscito, Napoli",
		AcceptLanguage: "it",
	})

	assert.Nil(t, err)
	// TODO assert result
}
