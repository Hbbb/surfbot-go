package surfline

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	queryString     = "resources=analysis&days=1&getAllSpots=false&units=e&usenearshore=true&interpolate=true&showOptimal=false"
	surflineSpotIds = []string{"2953", "2144", "131699"}
)

// Reports is a list of reports of reports
type Reports []Report

// Report represents the subset of the raw API response that we care about
type Report struct {
	ID        string   `json:"id"`
	Latitude  string   `json:"lat"`
	Longitude string   `json:"lon"`
	SpotName  string   `json:"name"`
	Surf      analysis `json:"Analysis"`
	MetaData  meta     `json:"_metadata"`
}

// Webpage returns the URL to the report on surfline.com
func (r Report) Webpage() string {
	spotPages := map[string]string{
		"South San Diego":        "http://www.surfline.com/surf-forecasts/southern-california/south-san-diego_2953/",
		"North San Diego":        "http://www.surfline.com/surf-forecasts/southern-california/north-san-diego_2144",
		"Nassau - Queens County": "http://www.surfline.com/surf-forecasts/long-island/nassau-queens-county_131699"}

	return spotPages[r.SpotName]
}

// analysis contains the specific surf report details
type analysis struct {
	Description []string `json:"generalText"`
	HeightText  []string `json:"surfRange"`
	MinHeight   []int    `json:"surfMin"`
	MaxHeight   []int    `json:"surfMax"`
}

func (a analysis) Text() string {
	return a.Description[0]
}

func (a analysis) Max() int {
	return a.MaxHeight[0]
}

func (a analysis) Min() int {
	return a.MinHeight[0]
}

type meta struct {
	CreatedAt string `json:"dateCreated"`
}

// GetReports fetches all surf reports in the spot id list
func GetReports() []Report {
	reports := Reports{}

	for _, id := range surflineSpotIds {
		report, err := fetchSurfReport(id)

		if err != nil {
			log.Fatal(err)
		}

		reports = append(reports, report)
	}

	return reports
}

func fetchSurfReport(id string) (Report, error) {
	report := Report{}
	url := fmt.Sprintf("http://api.surfline.com/v1/forecasts/%s?%s", id, queryString)
	res, err := http.Get(url)

	if err != nil {
		return Report{}, err
	}

	defer res.Body.Close()
	parsedReport, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return Report{}, err
	}

	json.Unmarshal(parsedReport, &report)

	return report, nil
}
