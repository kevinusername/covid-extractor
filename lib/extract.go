package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type row = struct {
	County    string
	Confirmed int64
	Deaths    int64
	Recovered int64
}

func checkHeader(h []string) bool {
	return h[0] == "FIPS" && h[7] == "Confirmed" && h[8] == "Deaths" && h[9] == "Recovered"
}

func countyData(fips, fileName string) (row, bool) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %s", fileName)
	}
	defer f.Close()

	r := csv.NewReader(f)
	record, err := r.Read()
	if err != nil || !checkHeader(record) {
		return row{}, false
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading csv file")
		}

		if record[0] == fips {
			confirmed, _ := strconv.ParseInt(record[7], 10, 64)
			deaths, _ := strconv.ParseInt(record[8], 10, 64)
			recovered, _ := strconv.ParseInt(record[9], 10, 64)

			cRow := row{
				County:    record[1],
				Confirmed: confirmed,
				Deaths:    deaths,
				Recovered: recovered,
			}
			return cRow, true
		}
	}
	return row{}, false
}