# Go Nominatim
A GO client for Nominatim API to perform both geocoding and reverse geocoding requests. No **api key** is needed.
The Nominatim API is free, but it has some limitation: be sure to take a look at the **Limitations** section.

## Installation
___
1. To install this Nominatim client for GO you need first to get the latest version of the package.
```
$ go get "github.com/doppiogancio/go-nominatim"
```

2. Import in your code
```
import nominatim "github.com/doppiogancio/go-nominatim"
```

## Quick Start
___

### Geocoding example
Converting the address "**Piazza del Plebiscito, Napoli**" in to a geographic coordinate with latitude and longitute.
```
package main

import nominatim "github.com/doppiogancio/go-nominatim"

func main() {
    coordinate, err := nominatim.Geocode("Piazza del Plebiscito, Napoli")    
}

```

Where the coordinate will look like:
```
{
    Latitude: 40.835855949999996,
    Longitude: 14.248565182098474,
}
```

### Reverse geocoding example
Reversing the geographic coordinate **(40.835855949999996, 14.248565182098474)** in to the address "Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia".
```
package main

import nominatim "github.com/doppiogancio/go-nominatim"

func main() {
    address, err := geocoder.ReverseGeocode(
        40.835855949999996,
        14.248565182098474,
        "it",
    )    
}

```

Where the address will look like:
```
{
  "DisplayName": "Piazza del Plebiscito, Via Cesario Console, San Ferdinando, Municipalità 1, Napoli, Campania, 80132, Italia",
  "Road": "Via Cesario Console",
  "Suburb": "San Ferdinando",
  "City": "Napoli",
  "County": "Napoli",
  "State": "Campania",
  "Postcode": "80132",
  "Country": "Italia",
  "CountryCode": "it"
}
```

## Nominatim API clients
If you prefer, you can use directly one of the 3 client to use the nominatim API.
1. **jsonv2** [README](geocoder/jsonv2/README.md)
2. **geojson** [README](geocoder/geojson/README.md)
3. **geocodejson** [README](geocoder/geocodejson/README.md)

## Limitations
https://operations.osmfoundation.org/policies/nominatim/

## API documentation
https://nominatim.org/release-docs/latest/
