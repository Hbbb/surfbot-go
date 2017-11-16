package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/hbbb/surfbot/slack"
	"github.com/hbbb/surfbot/surfline"
)

func main() {
	surfReports := surfline.GetReports()
	message := slack.BuildMessage(surfReports)
	payload, _ := json.Marshal(message)

	_, err := http.Post(slack.URL, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		log.Fatal(err)
	}
}
