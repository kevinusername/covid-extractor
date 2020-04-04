package county

import (
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

const DailyDataPath = "data/csse_covid_19_data/csse_covid_19_daily_reports/"

type Row struct {
	Updated   time.Time
	Confirmed int64
	Deaths    int64
	Recovered int64
}

type County struct {
	Name    string
	Records []Row
}

func (c *County) FromFiles(files []os.FileInfo) {
	c.Records = make([]Row, 0, 20)
	var wg sync.WaitGroup
	queue := make(chan Row, 1)

	for _, file := range files {
		if name := file.Name(); strings.HasSuffix(name, ".csv") {
			wg.Add(1)
			go func() {
				cData, ok := countyData(c.Name, DailyDataPath+name)
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
			c.Records = append(c.Records, r)
			wg.Done()
		}
	}()

	wg.Wait()
}

func (c *County) Sort() {
	sort.Slice(c.Records, func(i, j int) bool {
		return c.Records[i].Updated.After(c.Records[j].Updated)
	})
}
