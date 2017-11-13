package main

import (
	"github.com/hbbb/surfbot/slack"
	"github.com/hbbb/surfbot/surfline"
)

func main() {
	surfReports := surfline.GetReports()
	slack.SendMessage(surfReports)
}
