package shared

type (
	SearchRequest struct {
		Q              string
		Street         string
		City           string
		Country        string
		PostalCode     string
		AcceptLanguage string
	}

	ReverseGeocodeRequest struct {
		Latitude       float64
		Longitude      float64
		AcceptLanguage string
	}

	Coordinate struct {
		Latitude  float64
		Longitude float64
	}

	Address struct {
		DisplayName string
		Road        string
		Suburb      string
		City        string
		County      string
		State       string
		Postcode    string
		Country     string
		CountryCode string
	}
)
