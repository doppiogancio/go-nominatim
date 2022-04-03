package geojson

import (
	"fmt"
	"github.com/doppiogancio/go-nominatim/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Search(t *testing.T) {
	client := New("https://nominatim.openstreetmap.org")
	collection, err := client.Search(shared.SearchRequest{
		Q:              "Piazza del Plebiscito, Napoli",
		AcceptLanguage: "it",
	})

	assert.Nil(t, err)

	feature := collection.Features[0]
	assert.Equal(t, "Feature", feature.Type)
	assert.Equal(t, "Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia", feature.Properties.DisplayName)
}

func TestClient_Reverse(t *testing.T) {
	client := New("https://nominatim.openstreetmap.org")
	collection, err := client.Reverse(shared.ReverseGeocodeRequest{
		Latitude:  40.835855949999996,
		Longitude: 14.248565182098474,
	})

	assert.Nil(t, err)
	fmt.Println(collection)
	assert.Equal(
		t,
		"Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia",
		collection.Features[0].Properties.DisplayName,
	)
}
