package jsonv2

import (
	"github.com/doppiogancio/go-nominatim/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Search(t *testing.T) {
	client := New("https://nominatim.openstreetmap.org")
	places, err := client.Search(shared.SearchRequest{
		Q:              "Piazza del Plebiscito, Napoli",
		AcceptLanguage: "it",
	})

	assert.Nil(t, err)

	place := places[0]
	assert.Equal(t, "attraction", place.Type)
	assert.Equal(t, "Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia", place.DisplayName)
}

func TestClient_Reverse(t *testing.T) {
	client := New("https://nominatim.openstreetmap.org")
	place, err := client.Reverse(shared.ReverseGeocodeRequest{
		Latitude:  40.835855949999996,
		Longitude: 14.248565182098474,
	})

	assert.Nil(t, err)
	assert.Equal(
		t,
		"Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia",
		place.DisplayName,
	)
}
