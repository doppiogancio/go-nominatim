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
)
