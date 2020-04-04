package county

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"
)

func (c *County) WriteJSON() {
	jsonName := "out/json/" + strings.ReplaceAll(c.Name, " ", "") + ".json"
	outJSON, _ := os.OpenFile(jsonName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer outJSON.Close()
	enc := json.NewEncoder(outJSON)
	enc.Encode(c.Records)
}

func (c *County) WriteCSV() {
	csvName := "out/csv/" + strings.ReplaceAll(c.Name, " ", "") + ".csv"
	outCSV, _ := os.OpenFile(csvName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer outCSV.Close()
	w := csv.NewWriter(outCSV)
	w.Write([]string{"Updated", "Confirmed", "Deaths", "Recovered"})
	for _, record := range c.Records {
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
