package geocodejson

type (
	GeocodingProperties struct {
		PlaceID int    `json:"place_id"`
		OsmType string `json:"osm_type"`
		OsmID   int    `json:"osm_id"`
		Type    string `json:"type"`
		Label   string `json:"label"`
		Name    string `json:"name"`
	}

	Properties struct {
		Geocoding GeocodingProperties
	}

	Geometry struct {
		Type        string
		Coordinates []float64
	}

	Feature struct {
		Type       string
		Properties Properties
		Geometry   Geometry
	}

	Geocoding struct {
		Version     string
		Attribution string
		Licence     string
		Query       string
	}

	FeatureCollection struct {
		Type      string
		Geocoding Geocoding
		Features  []Feature
	}
)
