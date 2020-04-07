package main

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/kevinusername/SB-COVID-19/lib/county"
	"github.com/kevinusername/SB-COVID-19/lib/state"
)

var defaultCounties = []string{"Santa Barbara", "Los Angeles", "New York City"}

func oldmain() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = defaultCounties
	}

	files, err := ioutil.ReadDir(county.DailyDataPath)
	if err != nil {
		log.Fatal("Error reading data directory")
	}

	var wg sync.WaitGroup
	wg.Add(len(args))
	for _, cName := range args {
		go func(cName string) {
			c := county.County{Name: cName}
			c.FromFiles(files)
			c.Sort()

			c.WriteJSON()
			c.WriteCSV()

			wg.Done()
		}(cName)
	}
	wg.Wait()
}

var states = []string{"Alabama", "Alaska", "American Samoa", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "District of Columbia", "Federated States of Micronesia", "Florida", "Georgia", "Guam", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Marshall Islands", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Northern Mariana Islands", "Ohio", "Oklahoma", "Oregon", "Palau", "Pennsylvania", "Puerto Rico", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virgin Island", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}

func main() {
	sCapitaRecords := make([]state.CapitaRecord, 0, 50)
	for _, s := range states {
		stateRecord, ok := state.Extract(s, "04-06-2020.csv")
		if ok {
			c, d, ok := state.PerCapita(stateRecord)
			if ok {
				sCapitaRecords = append(sCapitaRecords, state.CapitaRecord{
					State:     s,
					Confirmed: c,
					Deaths:    d,
				})
			}
		}
	}
	state.Sort(sCapitaRecords)
	state.WriteCSV(sCapitaRecords)
}
