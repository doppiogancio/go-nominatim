package go_nominatim

import (
	"github.com/doppiogancio/go-nominatim/geocoder/geocodejson"
	"github.com/doppiogancio/go-nominatim/geocoder/geojson"
	"github.com/doppiogancio/go-nominatim/geocoder/jsonv2"
)

const (
	baseUrl = "https://nominatim.openstreetmap.org"
)

func NewJsonV2() *jsonv2.Client {
	return jsonv2.New(baseUrl)
}

func NewGeoJson() *geojson.Client {
	return geojson.New(baseUrl)
}

func NewGeocodeJson() *geocodejson.Client {
	return geocodejson.New(baseUrl)
}
