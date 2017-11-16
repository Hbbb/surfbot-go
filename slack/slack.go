package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hbbb/surfbot/surfline"
	"github.com/joho/godotenv"
)

// Slack API Format
// {
// 		 "title": "SurfBot"
//     "attachments": [
//         {
//             "color": "#36a64f",
//             "pretext": "Surf Reports 11/15/2017",
//             "title": "South San Diego",
//             "title_link": "https://surfline.com/",
//             "text": "3-4 ft. Occasional 5 ft. on standout sets",
//             "fields": [
//                 {
//                     "title": "Surf Max",
//                     "value": "5",
//                     "short": true
//                 },
// 				                {
//                     "title": "Surf Min",
//                     "value": "2",
//                     "short": true
//                 }
//             ]
//         }
//     ]
// }

type message struct {
	Title       string
	Attachments []attachment `json:"attachments"`
}

type attachment struct {
	Color       string       `json:"color"`
	Title       string       `json:"pretext"`
	Link        string       `json:"title_link"`
	Headline    string       `json:"text"`
	SurfHeights []surfHeight `json:"fields"`
}

type surfHeight struct {
	Title string `json:"title"`
	Value int    `json:"value"`
	Short bool   `json:"short"`
}

// SendMessage takes in a list of surf reports and sends them as a slack message
func SendMessage(surfReports []surfline.Report) {
	godotenv.Load()
	slackURL := os.Getenv("SLACK_URL")
	message := buildMessage(surfReports)

	payload, _ := json.Marshal(message)
	resp, err := http.Post(slackURL, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func buildMessage(reports []surfline.Report) attachment {
	message := message{Title: "Surf Report"}
	attachments := buildAttachments(reports)
	message.Attachments = attachments

	return attachment{}
}

func buildAttachments(reports []surfline.Report) []attachment {
	attachments := []attachment{}

	for _, report := range reports {
		attachment := attachment{
			Title:       report.SpotName,
			Color:       "#679AB0",
			Link:        fmt.Sprintf("https://new.surfline.com/"),
			Headline:    report.Surf.Text(),
			SurfHeights: buildFields(report),
		}

		attachments = append(attachments, attachment)
	}

	return attachments
}

func buildFields(r surfline.Report) []surfHeight {
	var fields []surfHeight

	maxHeight := surfHeight{
		Title: "Max Height",
		Value: r.Surf.Max(),
		Short: true}

	minHeight := surfHeight{
		Title: "Min Height",
		Value: r.Surf.Min(),
		Short: true}

	fields = append(fields, maxHeight)
	fields = append(fields, minHeight)

	return fields
}
