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
