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
    geocoder "github.com/doppiogancio/go-nominatim/geocoder/geojson"
    "github.com/doppiogancio/go-nominatim/shared"
)

func main() {
    request := shared.SearchRequest{
        Q:              "Piazza del Plebiscito, Napoli",
        AcceptLanguage: "it",
    }

    collection, err := geocoder.New().Search(request)
}
```

Where each **feature** of **collection.Features** will look like:

```
{
  "Type": "Feature",
  "Properties": {
    "PlaceId": 107822711,
    "OsmType": "way",
    "OsmId": 27182911,
    "DisplayName": "Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia",
    "PlaceRank": 30,
    "Category": "tourism",
    "Type": "attraction",
    "Importance": 0.6678657680745843,
    "Icon": "https://nominatim.openstreetmap.org/ui/mapicons//poi_point_of_interest.p.20.png"
  },
  "Bbox": [
    14.2474705,
    40.8350116,
    14.2498375,
    40.83674
  ],
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
https://nominatim.openstreetmap.org/reverse?format=geojson&lat=40.835855949999996&lon=14.248565182098474
```

An example:
```
package main

import (
    geocoder "github.com/doppiogancio/go-nominatim/geocoder/geojson"
    "github.com/doppiogancio/go-nominatim/shared"
)

func main() {
    request := shared.ReverseGeocodeRequest{
        Latitude:       40.835855949999996,
        Longitude:      14.248565182098474,
        AcceptLanguage: "it",
    }

    place, err := geocoder.New().Reverse(request)
}
```

Where the place will look like:
```
{
  "Type": "Feature",
  "Properties": {
    "PlaceID": 107822711,
    "OsmType": "way",
    "OsmID": 27182911,
    "DisplayName": "Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia",
    "PlaceRank": 30,
    "Category": "tourism",
    "Type": "attraction",
    "Importance": 0.36786576807458427,
    "Icon": null,
    "Addresstype": "tourism",
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
  },
  "Bbox": [
    14.2474705,
    40.8350116,
    14.2498375,
    40.83674
  ],
  "Geometry": {
    "Type": "Point",
    "Coordinates": [
      14.248565182098474,
      40.835855949999996
    ]
  }
}
```