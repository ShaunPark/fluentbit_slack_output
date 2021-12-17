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
	"reflect"
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
	sInfo := slackInfo{webhook: webhook, channel: channel, field: &field, title: &title, maxAttachment: maxAttatch}
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
	sInfo := output.FLBPluginGetContext(ctx).(slackInfo)
	log.Printf("[prettyslack] Flush called for webhook: %s", sInfo.channel)

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

		attachments = append(attachments, sInfo.makeAttachment(record))

		if len(attachments) == sInfo.maxAttachment {
			sendSlack(sInfo, attachments)
			attachments = []slack.Attachment{}
		}
	}

	if len(attachments) > 0 {
		sendSlack(sInfo, attachments)
	}

	return output.FLB_OK
}

const (
	DEFAULT_TITLE string = "Kernel logs by fluent-bit"
	BLOCK_TYPE    string = "header"
	TEXT_TYPE     string = "plain_text"
	EMPTY_STRING  string = ""
	DEFAULT_COLOR string = "#A9AAAA"
	UINT_ARR_TYPE string = "[]uint8"
	KEY_COLOR     string = "color"
)

func sendSlack(sInfo slackInfo, attachments []slack.Attachment) {
	blockType := BLOCK_TYPE
	textType := TEXT_TYPE
	title := DEFAULT_TITLE
	if sInfo.title != nil || *(sInfo.title) != EMPTY_STRING {
		title = *sInfo.title
	}
	emoji := true

	blocks := []slack.Block{}
	blocks = append(blocks, slack.Block{Type: &blockType, Text: &slack.Text{Type: &textType, Text: &title, Emoji: &emoji}})

	payload := slack.Payload{
		Attachments: attachments,
		Channel:     sInfo.channel,
		Blocks:      blocks,
	}

	err := slack.Send(sInfo.webhook, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}

func (s slackInfo) makeAttachment(data map[interface{}]interface{}) slack.Attachment {
	color := DEFAULT_COLOR
	msg := EMPTY_STRING
	attachment := slack.Attachment{Color: &color, Title: &msg}

	for key, val := range data {
		keyStr := key.(string)
		valStr := fmt.Sprintf("%v", val)
		if reflect.TypeOf(val).String() == UINT_ARR_TYPE {
			valStr = string(val.([]byte))
		}

		if keyStr == KEY_COLOR {
			attachment.Color = &valStr
		} else if *s.field == keyStr && keyStr != EMPTY_STRING {
			attachment.Title = &valStr
		} else {
			attachment.AddField(slack.Field{Title: keyStr, Value: valStr})
		}
	}
	return attachment
}

type slackInfo struct {
	webhook       string
	channel       string
	field         *string
	title         *string
	maxAttachment int
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
