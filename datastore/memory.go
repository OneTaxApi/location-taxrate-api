package datastore

import (
	"log"
	"os"
	"strings"
	"github.com/OneTaxApi/location-taxrate-api/loader"
)

type TaxRates struct {
	Store *[]*loader.TaxData `json:"store"`
}

func (b *TaxRates) Initialize() {
	filename := "./assets/locations.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	b.Store = loader.LoadData(file)
}

func (b *TaxRates) SearchLocation(iso string) *loader.TaxData {
	ret := Filter(b.Store, func(v *loader.TaxData) bool {
		return strings.ToLower(v.LocationIso) == strings.ToLower(iso)
	})
	if len(*ret) > 0 {
		return (*ret)[0]
	}
	return nil
}

func Filter(vs *[]*loader.TaxData, f func(*loader.TaxData) bool) *[]*loader.TaxData {
	vsf := make([]*loader.TaxData, 0)
	for _, v := range *vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return &vsf
}
