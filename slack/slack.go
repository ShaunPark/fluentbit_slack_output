package slack

import (
	"fmt"
	"log"

	"github.com/parnurzeal/gorequest"
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Action struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Url   string `json:"url"`
	Style string `json:"style"`
}

type Attachment struct {
	Fallback     *string   `json:"fallback,omitempty"`
	Color        *string   `json:"color,omitempty"`
	PreText      *string   `json:"pretext,omitempty"`
	AuthorName   *string   `json:"author_name,omitempty"`
	AuthorLink   *string   `json:"author_link,omitempty"`
	AuthorIcon   *string   `json:"author_icon,omitempty"`
	Title        *string   `json:"title,omitempty"`
	TitleLink    *string   `json:"title_link,omitempty"`
	Text         *string   `json:"text,omitempty"`
	ImageUrl     *string   `json:"image_url,omitempty"`
	Fields       []*Field  `json:"fields,omitempty"`
	Footer       *string   `json:"footer,omitempty"`
	FooterIcon   *string   `json:"footer_icon,omitempty"`
	Timestamp    *int64    `json:"ts,omitempty"`
	MarkdownIn   *[]string `json:"mrkdwn_in,omitempty"`
	Actions      []*Action `json:"actions,omitempty"`
	CallbackID   *string   `json:"callback_id,omitempty"`
	ThumbnailUrl *string   `json:"thumb_url,omitempty"`
	Blocks       []*Block  `json:"blocks,omitempty"`
}

type Payload struct {
	Parse       string       `json:"parse,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconUrl     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	LinkNames   string       `json:"link_names,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	UnfurlMedia bool         `json:"unfurl_media,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
	Blocks      []Block      `json:"blocks,omitempty"`
}

func (attachment *Attachment) AddField(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, &field)
	return attachment
}

func (attachment *Attachment) AddAction(action Action) *Attachment {
	attachment.Actions = append(attachment.Actions, &action)
	return attachment
}

func (attachment *Attachment) AddBlock(block Block) *Attachment {
	attachment.Blocks = append(attachment.Blocks, &block)
	return attachment
}

func redirectPolicyFunc(req gorequest.Request, via []gorequest.Request) error {
	return fmt.Errorf("incorrect token (redirection)")
}

func Send(webhookUrl string, proxy string, payload Payload) []error {
	request := gorequest.New().Proxy(proxy)
	request.Header.Set("Content-Type", "application/json")
	resp, _, err := request.
		Post(webhookUrl).
		RedirectPolicy(redirectPolicyFunc).
		Send(payload).
		End()

	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		log.Print(resp)
		return []error{fmt.Errorf("error sending msg. status: %v", resp.Status)}
	}

	return nil
}
