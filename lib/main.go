package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/kevinusername/SB-COVID-19/lib/county"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Must provide counties as arguements")
	}
	countyName := args[0]

	files, err := ioutil.ReadDir(county.DailyDataPath)
	if err != nil {
		log.Fatal("Error reading data directory")
	}

	c := county.County{Name: countyName}
	c.FromFiles(files)
	c.Sort()

	c.WriteJSON()
	c.WriteCSV()
}
