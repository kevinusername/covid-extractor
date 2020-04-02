package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	dailyDataPath    = "data/csse_covid_19_data/csse_covid_19_daily_reports/"
	santaBarbaraFIPS = "06083"
)

func main() {
	files, err := ioutil.ReadDir("data/csse_covid_19_data/csse_covid_19_daily_reports")
	if err != nil {
		log.Fatal("Error reading data directory")
	}

	countyRecords := make([]row, 0, 20)
	for _, file := range files {
		if name := file.Name(); strings.HasSuffix(name, ".csv") {
			cData, ok := countyData(santaBarbaraFIPS, dailyDataPath+name)
			if ok {
				countyRecords = append(countyRecords, cData)
			}
		}
	}
	fmt.Println(countyRecords)
}
