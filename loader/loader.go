package loader

import (
	"encoding/csv"
	"io"
	"log"
)

type TaxData struct {
	ISO     string  `json:"iso"`
	NAME    string  `json:"name"`
	RATE    string  `json:"rate"`
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
			ISO:    row[0],
			NAME:   row[1],
			RATE:   row[2],
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, taxrate)
	}
	return &ret
}
