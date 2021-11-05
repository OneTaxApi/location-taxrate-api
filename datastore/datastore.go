package datastore

import "github.com/OneTaxApi/location-taxrate-api/loader"

type TaxInfo interface {
	Initialize()
	SearchLocation(LocationIso string) *loader.TaxData
}
