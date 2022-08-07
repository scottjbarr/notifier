package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// SlackNotifier is used to send notifications to a Slack channel.
type SlackNotifier struct {
	URL         string
	ContentType string
	client      *http.Client
}

// NewSlackNotifier returns a new SlackNotifier for a given Slack channel.
func NewSlackNotifier(url string) SlackNotifier {
	return SlackNotifier{
		URL:         url,
		ContentType: "application/json",
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// Notify sends a simple message to the Slack channel.
//
// Example
//
//     curl -X POST -H 'content-type: application/json' --data '{"text":"Hello, World!"}' https://hooks.slack.com/services/xxx/yyy/zzzzzz
func (s SlackNotifier) Notify(b []byte) error {
	n := NewSlackNotification(string(b))

	b, err := json.Marshal(n)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(b)
	response, err := s.client.Post(s.URL, s.ContentType, reader)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("Received http status %v", response.StatusCode)
	}

	return nil
}

// SlackNotification is used to wrap a Slack notification.
type SlackNotification struct {
	Text string `json:"text"`
}

// NewSlackNotification returns a new Slack notification.
func NewSlackNotification(s string) SlackNotification {
	return SlackNotification{
		Text: s,
	}
}
