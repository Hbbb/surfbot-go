package slack

import (
	"os"
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/hbbb/surfbot/surfline"
)

// SendMessage takes in a list of surf reports and sends them as a slack message
func SendMessage(reports []surfline.Report) {
	godotenv.Load()
	slackURL := os.Getenv("SLACK_URL")

	payload, _ := json.Marshal(reports)
	http.Post(slackURL, "application/json", bytes.NewBuffer(payload))
}
