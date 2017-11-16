package slack

import (
	"os"

	"github.com/hbbb/surfbot/surfline"
	_ "github.com/joho/godotenv/autoload" // autoloads environment variables; remove before deploying
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

// URL is the slack webhook url
var URL = os.Getenv("SLACK_URL")

// Message is the top-level of the Slack webhook JSON body
type Message struct {
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

// BuildMessage converts a list of surfline Reports into Slack Messages
func BuildMessage(reports []surfline.Report) Message {
	message := Message{Title: "Surf Report"}
	attachments := buildAttachments(reports)
	message.Attachments = attachments

	return message
}

func buildAttachments(reports []surfline.Report) []attachment {
	attachments := []attachment{}

	for _, report := range reports {
		attachment := attachment{
			Title:       report.SpotName,
			Color:       "#679AB0",
			Link:        "https://new.surfline.com/", // TODO: Make this point to the current spot
			Headline:    report.Surf.Text(),
			SurfHeights: buildFields(report),
		}

		attachments = append(attachments, attachment)
	}

	return attachments
}

func buildFields(report surfline.Report) []surfHeight {
	var fields []surfHeight

	maxHeight := surfHeight{
		Title: "Max Height",
		Value: report.Surf.Max(),
		Short: true}

	minHeight := surfHeight{
		Title: "Min Height",
		Value: report.Surf.Min(),
		Short: true}

	fields = append(fields, maxHeight)
	fields = append(fields, minHeight)

	return fields
}
