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
	field := output.FLBPluginConfigKey(ctx, "textfield")
	format := output.FLBPluginConfigKey(ctx, "format")

	log.Printf("[prettyslack] webhook = %s#%s", webhook, channel)
	sInfo := slackInfo{webhook: webhook, channel: channel, field: &field, format: &format}
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

		switch *sInfo.format {
		case "kernel":
			attachments = append(attachments, sInfo.makeKernelAttachment(record))
		default:
			attachments = append(attachments, makeJsonAttachment(record))
		}
	}

	header := "header"
	pText := "plain_text"
	hdrMsg := "Kernel logs by fluent-bit"
	emoji := true

	blocks := []slack.Block{}
	blocks = append(blocks, slack.Block{Type: &header, Text: &slack.TextBlock{Type: &pText, Text: &hdrMsg, Emoji: &emoji}})

	payload := slack.Payload{
		Attachments: attachments,
		Channel:     sInfo.channel,
		Blocks:      blocks,
	}

	err := slack.Send(sInfo.webhook, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}

	return output.FLB_OK
}

func (s slackInfo) makeKernelAttachment(data map[interface{}]interface{}) slack.Attachment {
	var color, msg *string
	fields := []*slack.Field{}
	for key, val := range data {
		log.Printf("%s, %v", key, val)
	}

	for key, val := range data {
		keyStr := key.(string)
		if "color" == "keyStr" {
			color = val.(*string)
		} else if *s.field == keyStr {
			log.Printf("%s %s", *s.field, keyStr)
			*msg = fmt.Sprintf("%s", val)
		} else {
			switch val.(type) {
			case []uint8:
				fields = append(fields, &slack.Field{Title: keyStr, Value: fmt.Sprintf("%s", val)})
			default:
				fields = append(fields, &slack.Field{Title: keyStr, Value: fmt.Sprintf("%v", val)})
			}
		}
	}

	if color == nil {
		*color = "#A9AAAA"
	}
	if msg == nil {
		*msg = ""
	}
	attachment1 := slack.Attachment{Color: color, Text: msg, Fields: fields}
	return attachment1
}

func makeJsonAttachment(data map[interface{}]interface{}) slack.Attachment {
	return slack.Attachment{}
}

type slackInfo struct {
	webhook string
	channel string
	field   *string
	format  *string
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
