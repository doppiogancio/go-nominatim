package jsonv2

type (
	Place struct {
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
)
