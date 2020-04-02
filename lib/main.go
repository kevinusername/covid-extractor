package main

import "fmt"

const (
	dailyDataPath    = "data/csse_covid_19_data/csse_covid_19_daily_reports/"
	santaBarbaraFIPS = "06083"
)

func main() {
	sbData, _ := countyData(santaBarbaraFIPS, dailyDataPath+"04-01-2020.csv")
	fmt.Printf("%v", sbData)
}
