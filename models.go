package go_nominatim

type (
	Location struct {
		PlaceId     int `json:"place_id"`
		Licence     string
		OsmType     string `json:"osm_type"`
		OsmId       int    `json:"osm_id"`
		Lat         string
		Lon         string
		DisplayName string `json:"display_name"`
		PlaceRank   int    `json:"place_rank"`
		Category    string
		Type        string
		Importance  float64
		Icon        string
	}

	SearchRequest struct {
		Q              string
		Street         string
		City           string
		Country        string
		PostalCode     string
		Format         Format
		AcceptLanguage string
	}

	Format string
)

const (
	FORMAT_XML          Format = "xml"
	FORMAT_JSON         Format = "json"
	FORMAT_JSONv2       Format = "jsonv2"
	FORMAT_GEO_JSON     Format = "geojson"
	FORMAT_GEOCODE_JSON Format = "geocodejson"
)
