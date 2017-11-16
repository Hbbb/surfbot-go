package slack

import (
	"os"

	"github.com/hbbb/surfbot/surfline"
	_ "github.com/joho/godotenv/autoload" // autoloads environment variables; remove before deploying
)

// Slack Message Format
// https://api.slack.com/docs/messages/builder?msg=%7B%22attachments%22%3A%5B%7B%22color%22%3A%22%2336a64f%22%2C%22pretext%22%3A%22Surf%20Reports%2011%2F15%2F2017%22%2C%22title%22%3A%22South%20San%20Diego%22%2C%22title_link%22%3A%22https%3A%2F%2Fsurfline.com%2F%22%2C%22text%22%3A%223-4%20ft.%20Occasional%205%20ft.%20on%20standout%20sets%22%2C%22fields%22%3A%5B%7B%22title%22%3A%22Surf%20Max%22%2C%22value%22%3A%225%22%2C%22short%22%3Atrue%7D%2C%7B%22title%22%3A%22Surf%20Min%22%2C%22value%22%3A%222%22%2C%22short%22%3Atrue%7D%5D%7D%5D%7D

// URL is the slack webhook url
var URL = os.Getenv("SLACK_URL")

// Message is the top-level of the Slack webhook JSON body
type Message struct {
	Title       string
	Attachments []attachment `json:"attachments"`
}

type attachment struct {
	Color       string       `json:"color"`
	Title       string       `json:"title"`
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
			Link:        report.Webpage(),
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
