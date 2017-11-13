package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	surfineBaseURL      = "http://api.surfline.com/v1/forecasts/"
	surflineQueryParams = "?resources=analysis&days=1&getAllSpots=false&units=e&usenearshore=true&interpolate=true&showOptimal=false"
)

var surflineSpotIds = []string{"2953", "2144", "131699"}

func main() {
	reports := GetReports()
	fmt.Println("Surf Report", reports)
}

// GetReports fetches all surf reports in the spot id list
func GetReports() []Report {
	reports := []Report{}

	for _, id := range surflineSpotIds {
		report, err := fetchSurfReport(id)

		if err != nil {
			log.Fatal(err)
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	return reports
}

func fetchSurfReport(id string) (Report, error) {
	report := Report{}
	url := surfineBaseURL + id + surflineQueryParams
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		return Report{}, err
	}

	parsedReport, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
		return Report{}, err
	}

	json.Unmarshal(parsedReport, &report)

	defer res.Body.Close()

	return report, nil
}

// analysis contains the specific surf report details
type analysis struct {
	Description []string `json:"generalText"`
	HeightText  []string `json:"surfRange"`
	SurfMin     []int `json:"surfMin"`
	SurfMax     []int `json:"surfMax"`
}

type meta struct {
	CreatedAt string `json:"dateCreated"`
}

// Report represents the subset of the raw API response that we care about
type Report struct {
	ID        string   `json:"id"`
	Latitude  string   `json:"lat"`
	Longitude string   `json:"lon"`
	SpotName  string   `json:"name"`
	Surf      analysis `json:"Analysis"`
	// MetaData  meta     `json:"_metadata"`
}
