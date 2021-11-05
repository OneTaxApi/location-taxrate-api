package loader

import (
	"encoding/csv"
	"io"
	"log"
)

type TaxData struct {
	LocationIso     string  `json:"location_iso"`
	LocationName    string  `json:"location_name"`
	TaxRate         string  `json:"location_taxrate"`
}

func LoadData(r io.Reader) *[]*TaxData {
	reader := csv.NewReader(r)

	ret := make([]*TaxData, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		if err != nil {
			log.Println(err)
		}
		taxrate := &TaxData{
			LocationIso:    row[0],
			LocationName:   row[1],
			TaxRate:        row[2],
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, taxrate)
	}
	return &ret
}
