package state

import (
	"sort"
	"time"
)

type Record struct {
	State     string
	Confirmed int
	Deaths    int
	Updated   time.Time
}

type CapitaRecord struct {
	State     string
	Confirmed float64
	Deaths    float64
}

func Sort(cr []CapitaRecord) {
	sort.Slice(cr, func(i, j int) bool {
		return cr[i].Confirmed > cr[j].Confirmed
	})
}
