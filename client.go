package go_nominatim

import (
	"errors"
	"github.com/doppiogancio/go-nominatim/geocoder/geocodejson"
	"github.com/doppiogancio/go-nominatim/geocoder/geojson"
	"github.com/doppiogancio/go-nominatim/geocoder/jsonv2"
	"github.com/doppiogancio/go-nominatim/shared"
	"strconv"
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

func Geocode(address string) (*shared.Coordinate, error) {
	locations, err := NewJsonV2().Search(shared.SearchRequest{
		Q: address,
	})
	if err != nil {
		return nil, err
	}

	if len(locations) == 0 {
		return nil, errors.New("address not found")
	}

	latitude, err := strconv.ParseFloat(locations[0].Lat, 64)
	if err != nil {
		return nil, err
	}

	longitude, err := strconv.ParseFloat(locations[0].Lon, 64)
	if err != nil {
		return nil, err
	}

	return &shared.Coordinate{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil

}

func ReverseGeocode(latitude, longitude float64, language string) (*shared.Address, error) {
	place, err := NewJsonV2().Reverse(shared.ReverseGeocodeRequest{
		Latitude:       latitude,
		Longitude:      longitude,
		AcceptLanguage: language,
	})

	if err != nil {
		return nil, err
	}

	return &shared.Address{
		DisplayName: place.DisplayName,
		Road:        place.Address.Road,
		Suburb:      place.Address.Suburb,
		City:        place.Address.City,
		County:      place.Address.County,
		State:       place.Address.State,
		Postcode:    place.Address.Postcode,
		Country:     place.Address.Country,
		CountryCode: place.Address.CountryCode,
	}, nil
}
