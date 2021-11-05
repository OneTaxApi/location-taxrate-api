package datastore

import "github.com/OneTaxApi/location-taxrate-api/loader"

type TaxInfo interface {
	Initialize()
	SearchLocation(LocationIso, LocationName string, taxRate float64) *[]*loader.TaxData
}
