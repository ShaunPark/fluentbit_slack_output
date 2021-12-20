package main

import (
	"C"
	"log"
	"unsafe"

	"github.com/ShaunPark/fluentbit_slack_output/slack"
	"github.com/fluent/fluent-bit-go/output"
)
import (
	"fmt"
	"strconv"
	"time"
)

//export FLBPluginRegister
func FLBPluginRegister(ctx unsafe.Pointer) int {
	// Gets called only once when the plugin.so is loaded
	log.Print("[prettyslack] register")

	return output.FLBPluginRegister(ctx, "prettyslack", "Slack output pretty")
}

//export FLBPluginInit
func FLBPluginInit(ctx unsafe.Pointer) int {
	webhook := output.FLBPluginConfigKey(ctx, "webhook")
	channel := output.FLBPluginConfigKey(ctx, "channel")
	field := output.FLBPluginConfigKey(ctx, "message_field")
	title := output.FLBPluginConfigKey(ctx, "title")
	maxAttachment := output.FLBPluginConfigKey(ctx, "max_attachment")

	maxAttatch, err := strconv.Atoi(maxAttachment)
	if err != nil {
		maxAttatch = 5
	}

	if maxAttatch > 20 {
		maxAttatch = 20
	}

	log.Printf("[prettyslack] webhook = %s#%s", webhook, channel)
	sInfo := slack.SlackInfo{Webhook: webhook, Channel: channel, Field: &field, Title: &title, MaxAttachment: maxAttatch}
	// Set the context to point to any Go variable
	output.FLBPluginSetContext(ctx, sInfo)
	return output.FLB_OK
}

func FLBPluginFlush(data unsafe.Pointer, length C.int, tag *C.char) int {
	log.Print("[prettyslack] Flush called for unknown instance")
	return output.FLB_OK
}

//export FLBPluginFlushCtx
func FLBPluginFlushCtx(ctx, data unsafe.Pointer, length C.int, tag *C.char) int {
	// Gets called with a batch of records to be written to an instance.
	sInfo := output.FLBPluginGetContext(ctx).(slack.SlackInfo)
	log.Printf("[prettyslack] Flush called for webhook: %s", sInfo.Channel)

	dec := output.NewDecoder(data, int(length))

	attachments := []slack.Attachment{}

	for {
		ret, ts, record := output.GetRecord(dec)
		if ret != 0 {
			break
		}

		var timestamp time.Time
		switch t := ts.(type) {
		case output.FLBTime:
			timestamp = ts.(output.FLBTime).Time
		case uint64:
			timestamp = time.Unix(int64(t), 0)
		default:
			fmt.Println("time provided invalid, defaulting to now.")
			timestamp = time.Now()
		}
		record["timestamp"] = timestamp

		attachments = append(attachments, sInfo.MakeAttachment(record))

		if len(attachments) == sInfo.MaxAttachment {
			sInfo.SendSlack(attachments)
			attachments = []slack.Attachment{}
		}
	}

	if len(attachments) > 0 {
		sInfo.SendSlack(attachments)
	}

	return output.FLB_OK
}

//export FLBPluginExit
func FLBPluginExit() int {
	return output.FLB_OK
}

//export FLBPluginExitCtx
func FLBPluginExitCtx(ctx unsafe.Pointer) int {
	return output.FLB_OK
}

func main() {
}
