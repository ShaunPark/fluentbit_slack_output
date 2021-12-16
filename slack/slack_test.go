package slack_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/ShaunPark/fluentbit_slack_output/slack"
	"github.com/parnurzeal/gorequest"
)

const webhook = "https://hooks.slack.com/services/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"
const channel = "#fluentbit_test"

func TestPtrn(t *testing.T) {
	// sInfo := slack.NewSlackInfo(webhook, channel, nil)
	color := "#FF336A"
	msg := "Disk error"
	attachment1 := slack.Attachment{Color: &color, Text: &msg}
	attachment1.AddField(slack.Field{Title: "Author", Value: "Ashwanth Kumar"}).AddField(slack.Field{Title: "Status", Value: "Completed"})
	attachment2 := slack.Attachment{Color: &color, Text: &msg}
	attachment2.AddField(slack.Field{Title: "Author", Value: "Ashwanth Kumar"}).AddField(slack.Field{Title: "Status", Value: "Completed"})

	blocks := []slack.Block{}

	header := "header"
	pText := "plain_text"
	context := "context"
	emoji := true
	mrkdwn := "mrkdwn"
	msg1 := "*This* is :smile: markdown"
	msg2 := "Author: K A Applegate"
	hdrMsg := "Header Message"
	blocks = append(blocks, slack.Block{Type: &header, Text: &slack.TextBlock{Type: &pText, Text: &hdrMsg, Emoji: &emoji}})

	elements := []slack.Element{}
	elements = append(elements, slack.Element{Type: &mrkdwn, Text: &msg1})
	elements = append(elements, slack.Element{Type: &pText, Text: &msg2, Emoji: &emoji})

	blocks = append(blocks, slack.Block{Type: &context, Elements: &elements})
	payload := slack.Payload{
		// Text: "test",
		// Channel: channel,
		Blocks:      blocks,
		Attachments: []slack.Attachment{attachment1, attachment2},
	}
	bytes, _ := json.Marshal(payload)
	log.Print(string(bytes))

	//test := "{\"blocks\": [{\"type\": \"section\",\"text\": {\"type\": \"plain_text\",\"text\": \"This is a plain text section block.\",\"emoji\": true}}]}"

	err := Send(webhook, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}

func redirectPolicyFunc(req gorequest.Request, via []gorequest.Request) error {
	return fmt.Errorf("incorrect token (redirection)")
}

func Send(webhookUrl string, proxy string, payload slack.Payload) []error {
	request := gorequest.New().Proxy(proxy)
	resp, _, err := request.
		Post(webhookUrl).
		RedirectPolicy(redirectPolicyFunc).
		Send(payload).
		End()

	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return []error{fmt.Errorf("error sending msg. status: %v", resp.Status)}
	}

	return nil
}
