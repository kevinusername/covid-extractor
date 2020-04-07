package state

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const dailyDataPath = "data/csse_covid_19_data/csse_covid_19_daily_reports/"
const datePattern = "01-02-2006"

var headerRow = []string{"FIPS", "Admin2", "Province_State", "Country_Region", "Last_Update", "Lat", "Long_", "Confirmed", "Deaths", "Recovered", "Active", "Combined_Key"}

func checkHeader(row []string) bool {
	for i, v := range row {
		if headerRow[i] != v {
			return false
		}
	}
	return true
}

func parseRecord(record []string) (Record, bool) {
	confirmed, e1 := strconv.ParseInt(record[7], 10, 0)
	deaths, e2 := strconv.ParseInt(record[8], 10, 0)
	if e1 != nil || e2 != nil {
		return Record{}, false
	}

	sRow := Record{
		Confirmed: int(confirmed),
		Deaths:    int(deaths),
	}
	return sRow, true

}

func fileDate(fName string) time.Time {
	dateString := strings.TrimSuffix(fName, ".csv")
	d, err := time.Parse(datePattern, dateString)
	if err != nil {
		panic("Failed parsing date from file name")
	}
	return d
}

func Extract(stateName, fileName string) (Record, bool) {
	f, err := os.Open(dailyDataPath + fileName)
	if err != nil {
		log.Fatalf("Error opening file: %s", fileName)
	}
	defer f.Close()

	r := csv.NewReader(f)
	header, err := r.Read()
	if err != nil || !checkHeader(header) {
		return Record{}, false
	}

	sRecord := Record{
		State:   stateName,
		Updated: fileDate(fileName),
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading csv file")
		}

		if record[2] == stateName {
			sRow, ok := parseRecord(record)
			if ok {
				sRecord.Confirmed += sRow.Confirmed
				sRecord.Deaths += sRow.Deaths
			}
		}
	}
	return sRecord, true
}
