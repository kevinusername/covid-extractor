package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	dailyDataPath = "data/csse_covid_19_data/csse_covid_19_daily_reports/"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Must provide counties as arguements")
	}
	county := args[0]

	files, err := ioutil.ReadDir(dailyDataPath)
	if err != nil {
		log.Fatal("Error reading data directory")
	}

	countyRecords := make([]row, 0, 20)
	var wg sync.WaitGroup
	queue := make(chan row, 1)

	for _, file := range files {
		if name := file.Name(); strings.HasSuffix(name, ".csv") {
			wg.Add(1)
			go func() {
				cData, ok := countyData(county, dailyDataPath+name)
				if ok {
					queue <- cData
				} else {
					wg.Done()
				}
			}()
		}
	}

	go func() {
		for r := range queue {
			countyRecords = append(countyRecords, r)
			wg.Done()
		}
	}()

	wg.Wait()

	sort.Slice(countyRecords, func(i, j int) bool {
		return countyRecords[i].Updated.After(countyRecords[j].Updated)
	})

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
