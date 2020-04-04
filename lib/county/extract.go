package county

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const datePattern = "2006-01-02 15:04:05"

func checkHeader(h []string) bool {
	return h[1] == "Admin2" && h[7] == "Confirmed" && h[8] == "Deaths" && h[9] == "Recovered"
}

func parseRecord(record []string) (Row, bool) {
	confirmed, e1 := strconv.ParseInt(record[7], 10, 64)
	deaths, e2 := strconv.ParseInt(record[8], 10, 64)
	recovered, e3 := strconv.ParseInt(record[9], 10, 64)
	updated, e4 := time.Parse(datePattern, record[4])
	if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
		return Row{}, false
	}

	updated = updated.In(time.Local)
	cRow := Row{Updated: updated, Confirmed: confirmed, Deaths: deaths, Recovered: recovered}
	return cRow, true

}

func countyData(countyName, fileName string) (Row, bool) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %s", fileName)
	}
	defer f.Close()

	r := csv.NewReader(f)
	record, err := r.Read()
	if err != nil || !checkHeader(record) {
		return Row{}, false
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading csv file")
		}

		if record[1] == countyName {
			cRow, ok := parseRecord(record)
			if ok {
				return cRow, true
			}
		}
	}
	return Row{}, false
}
