package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	dailyDataPath = "data/csse_covid_19_data/csse_covid_19_daily_reports/"
	county        = "Santa Barbara"
)

func main() {
	files, err := ioutil.ReadDir(dailyDataPath)
	if err != nil {
		log.Fatal("Error reading data directory")
	}

	countyRecords := make([]row, 0, 20)
	for _, file := range files {
		if name := file.Name(); strings.HasSuffix(name, ".csv") {
			cData, ok := countyData(county, dailyDataPath+name)
			if ok {
				countyRecords = append(countyRecords, cData)
			}
		}
	}

	jsonName := "out/json/" + strings.ReplaceAll(county, " ", "") + ".json"
	outJSON, _ := os.OpenFile(jsonName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	defer outJSON.Close()
	enc := json.NewEncoder(outJSON)
	enc.Encode(countyRecords)

	csvName := "out/csv/" + strings.ReplaceAll(county, " ", "") + ".csv"
	outCSV, _ := os.OpenFile(csvName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	defer outCSV.Close()
	w := csv.NewWriter(outCSV)
	w.Write([]string{"Updated", "Confirmed", "Deaths", "Recovered"})
	for _, record := range countyRecords {
		sRecord := []string{
			record.Updated.Format(time.RFC822),
			strconv.FormatInt(record.Confirmed, 10),
			strconv.FormatInt(record.Deaths, 10),
			strconv.FormatInt(record.Recovered, 10),
		}
		w.Write(sRecord)
	}
	w.Flush()
}
