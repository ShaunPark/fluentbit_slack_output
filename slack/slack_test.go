package slack_test

import (
	"testing"

	"github.com/ShaunPark/fluentbit_slack_output/slack"
)

func TestPtrn(t *testing.T) {
	fName := "msg"
	title := "Title message"
	sInfo := slack.SlackInfo{Field: &fName, Webhook: "https://hooks.slack.com/services/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", Channel: "fluent-test", Title: &title}
	attachments := []slack.Attachment{}
	records := make(map[interface{}]interface{})

	records["color"] = "#34cceb"
	records["msg"] = "TEST MESSAGE"
	records["nodeName"] = "nodeName"
	attachments = append(attachments, sInfo.MakeAttachment(records))
	sInfo.SendSlack(attachments)
	// Test()
}
