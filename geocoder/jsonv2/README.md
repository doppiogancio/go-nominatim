# Geocode JSON Client

## Search Request
The code below will submit a request like:
```
https://nominatim.openstreetmap.org/search.php?q=Piazza del Plebiscito, Napoli&format=jsonv2
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

    places, err := geocoder.NewJsonV2().Search(request)
}
```

Where each place will look like:

```
{
    PlaceID: 107822711,
    Licence: "Data © OpenStreetMap contributors, ODbL 1.0. https://osm.org/copyright",
    OsmType: "way",
    OsmID: 27182911,
    Lat: "40.835855949999996",
    Lon: "14.248565182098474",
    DisplayName: "Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia",
    PlaceRank: 30,
    Category: "tourism",
    Type: "attraction",
    Importance: 0.7778657680745842,
    Icon: "https://nominatim.openstreetmap.org/ui/mapicons//poi_point_of_interest.p.20.png"
}
```

## Reverse Request
The code below will submit a request like:
```
https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=40.835855949999996&lon=14.248565182098474
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

    place, err := geocoder.NewJsonV2().Reverse(request)
}
```

Where the place will look like:
```
{
  "PlaceID": 107822711,
  "Licence": "Data © OpenStreetMap contributors, ODbL 1.0. https://osm.org/copyright",
  "OsmType": "way",
  "OsmID": 27182911,
  "Lat": "40.835855949999996",
  "Lon": "14.248565182098474",
  "DisplayName": "Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia",
  "PlaceRank": 30,
  "Category": "tourism",
  "Type": "attraction",
  "Importance": 0.36786576807458427,
  "Icon": "",
  "addresstype": "tourism",
  "Name": "Piazza del Plebiscito",
  "Address": {
    "Tourism": "Piazza del Plebiscito",
    "Road": "Via Cesario Console",
    "Suburb": "San Ferdinando",
    "City": "Napoli",
    "County": "Napoli",
    "State": "Campania",
    "Postcode": "80132",
    "Country": "Italia",
    "country_code": "it"
  }
}
```