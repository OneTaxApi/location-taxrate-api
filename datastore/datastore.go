package datastore

import "github.com/OneTaxApi/location-taxrate-api/location-taxrate-api/api/loader"

type TaxInfo interface {
	Initialize()
	SearchLocation(locationIso, locationName string, taxrate float64) *[]*loader.TaxData
}
