# Go Nominatim
A GO client for Nominatim API

## Example
### Json v2
#### Request
```
client := NewJsonV2()

places, err := client.Search(shared.SearchRequest{
    Q:              "Piazza del Plebiscito, Napoli",
    AcceptLanguage: "it",
})

```

#### Response
```
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
```

### Geo Json
#### Request
```
client := NewGeoJson()

collection, err := client.Search(shared.SearchRequest{
    Q:              "Piazza del Plebiscito, Napoli",
    AcceptLanguage: "it",
})
```

#### Response
```
Properties struct {
    PlaceID     int    `json:"place_id"`
    OsmType     string `json:"osm_type"`
    OsmID       int    `json:"osm_id"`
    DisplayName string `json:"display_name"`
    PlaceRank   int    `json:"place_rank"`
    Category    string
    Type        string
    Importance  float64
    Icon        string
}

Geometry struct {
    Type        string
    Coordinates []float64
}

Feature struct {
    Type       string
    Properties Properties
    Bbox       []float64
    Geometry   Geometry
}

FeatureCollection struct {
    Type     string
    Licence  string
    Features []Feature
}
```

### Geocode Json
#### Request
```
client := NewGeocodeJson()

collection, err := client.Search(shared.SearchRequest{
    Q:              "Piazza del Plebiscito, Napoli",
    AcceptLanguage: "it",
})
```

#### Response
```
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
```


## API documentation
https://nominatim.org/release-docs/latest/api/Search/
