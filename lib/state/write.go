package state

import (
	"encoding/csv"
	"os"
	"strconv"
)

func WriteCSV(records []CapitaRecord) {
	csvName := "out/state-capita.csv"
	outCSV, _ := os.OpenFile(csvName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer outCSV.Close()
	w := csv.NewWriter(outCSV)
	w.Write([]string{"State", "Confirmed (per-100k)", "Deaths (per-100k)"})
	for _, record := range records {
		sRecord := []string{
			record.State,
			strconv.FormatFloat(record.Confirmed, 'f', 2, 64),
			strconv.FormatFloat(record.Deaths, 'f', 2, 64),
		}
		w.Write(sRecord)
	}
	w.Flush()
}
