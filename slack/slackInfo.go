package slack

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
)

const (
	DEFAULT_TITLE      string = "Kernel logs by fluent-bit"
	BLOCK_TYPE         string = "header"
	TEXT_TYPE          string = "plain_text"
	EMPTY_STRING       string = ""
	DEFAULT_COLOR      string = "#A9AAAA"
	UINT_ARR_TYPE      string = "[]uint8"
	KEY_COLOR          string = "color"
	BLOCK_TYPE_SECTION string = "section"
	TEXT_TYPE_MRKDWN   string = "mrkdwn"
)

type SlackInfo struct {
	Webhook       string
	Channel       string
	Field         *string
	Title         *string
	MaxAttachment int
}

func (s *SlackInfo) SendSlack(attachments []Attachment) {
	blockType := BLOCK_TYPE
	textType := TEXT_TYPE
	title := DEFAULT_TITLE
	if s.Title != nil || *(s.Title) != EMPTY_STRING {
		title = *s.Title
	}
	emoji := true

	blocks := []Block{}
	blocks = append(blocks, Block{Type: &blockType, Text: &Text{Type: &textType, Text: &title, Emoji: &emoji}})

	payload := Payload{
		Attachments: attachments,
		Channel:     s.Channel,
		Blocks:      blocks,
	}
	bytes, _ := json.Marshal(payload)
	log.Print(string(bytes))
	err := Send(s.Webhook, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}

func (s *SlackInfo) MakeAttachment(data map[interface{}]interface{}) Attachment {
	color := DEFAULT_COLOR
	msg := "title message"
	attachment := Attachment{Color: &color}

	fieldStrs := []string{}
	for key, val := range data {
		keyStr := key.(string)
		valStr := fmt.Sprintf("%v", val)
		if reflect.TypeOf(val).String() == UINT_ARR_TYPE {
			valStr = string(val.([]byte))
		}

		if keyStr == KEY_COLOR {
			attachment.Color = &valStr
		} else if *s.Field == keyStr && keyStr != EMPTY_STRING {
			msg = valStr
		} else {
			fieldStrs = append(fieldStrs, fmt.Sprintf("%s: `%s`", strings.ToUpper(keyStr[:1])+keyStr[1:], strings.TrimSpace(valStr)))
		}
	}
	blockType := BLOCK_TYPE_SECTION
	textType := TEXT_TYPE_MRKDWN
	sort.Strings(fieldStrs)
	text := strings.Join(fieldStrs, "\n")
	log.Print(text)
	msg = fmt.Sprintf("*%s*", msg)
	attachment.AddBlock(Block{Type: &blockType, Text: &Text{Type: &textType, Text: &msg}})
	attachment.AddBlock(Block{Type: &blockType, Text: &Text{Type: &textType, Text: &text}})
	return attachment
}
