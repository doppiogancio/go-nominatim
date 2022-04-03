# Geocode JSON Client

## Search Request
The code below will submit a request like:
```
https://nominatim.openstreetmap.org/search.php?q=Piazza del Plebiscito, Napoli&format=geocodejson
```

An example:
```
package main

import (
    geocoder "github.com/doppiogancio/go-nominatim"
    "github.com/doppiogancio/go-nominatim/shared"
)

func main() {
    request := shared.SearchRequest{
        Q:              "Piazza del Plebiscito, Napoli",
        AcceptLanguage: "it",
    }

    places, err := geocoder.NewGeocodeJson().Search(request)
}
```

Where each place will look like:

```
{
  "Type": "Feature",
  "Properties": {
    "Geocoding": {
      "PlaceID": 107822711,
      "OsmType": "way",
      "OsmID": 27182911,
      "Type": "attraction",
      "Label": "Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia",
      "Name": "Piazza del Plebiscito"
    }
  },
  "Geometry": {
    "Type": "Point",
    "Coordinates": [
      14.248565182098474,
      40.835855949999996
    ]
  }
}
```

## Reverse Request
The code below will submit a request like:
```
https://nominatim.openstreetmap.org/reverse?format=geocodejson&lat=40.835855949999996&lon=14.248565182098474
```

An example:
```
package main

import (
    geocoder "github.com/doppiogancio/go-nominatim"
    "github.com/doppiogancio/go-nominatim/shared"
)

func main() {
    request := shared.ReverseGeocodeRequest{
        Latitude:       40.835855949999996,
        Longitude:      14.248565182098474,
        AcceptLanguage: "it",
    }

    place, err := geocoder.NewGeocodeJson().Reverse(request)
}
```

Where the place will look like:
```
not yet implemented
```