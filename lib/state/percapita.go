package state

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func getPopulation(state string) (int64, bool) {
	f, err := os.Open("state-pop.csv")
	if err != nil {
		log.Fatal("Error opening state-pop.csv")
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, _ := r.ReadAll()

	for _, rec := range records {
		if rec[0] == state {
			pop, err := strconv.ParseInt(rec[1], 10, 0)
			if err != nil {
				log.Fatal("Error reading state population")
			}
			return pop, true
		}
	}
	return 0, false
}

func PerCapita(r Record) (float64, float64, bool) {
	population, ok := getPopulation(r.State)
	if !ok {
		return 0, 0, false
	}
	hThousands := float64(population) / 100000
	return float64(r.Confirmed) / hThousands, float64(r.Deaths) / hThousands, true
}
