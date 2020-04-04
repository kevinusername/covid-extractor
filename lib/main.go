package main

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/kevinusername/SB-COVID-19/lib/county"
)

var defaultCounties = []string{"Santa Barbara", "Los Angeles", "New York City"}

func main() {
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
