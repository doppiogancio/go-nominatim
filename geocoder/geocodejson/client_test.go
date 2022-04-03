package geocodejson

import (
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
	assert.Equal(t, "Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia", feature.Properties.Geocoding.Label)
	assert.Equal(t, "Piazza del Plebiscito", feature.Properties.Geocoding.Name)

}
