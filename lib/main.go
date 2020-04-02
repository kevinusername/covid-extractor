package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

	outFileName := "out/" + strings.ReplaceAll(county, " ", "") + ".json"
	outFile, _ := os.OpenFile(outFileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	defer outFile.Close()
	enc := json.NewEncoder(outFile)

	enc.Encode(countyRecords)
}
