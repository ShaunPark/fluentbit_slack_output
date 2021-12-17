package slack

type Text struct {
	Type  *string `json:"type"`
	Text  *string `json:"text,omitempty"`
	Emoji *bool   `json:"emoji,omitempty"`
}

type Blocks struct {
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Type      *string    `json:"type"`
	Text      *Text      `json:"text,omitempty"`
	Title     *Text      `json:"title,omitempty"`
	Fields    *[]Text    `json:"fields,omitempty"`
	Accessory *Accessory `json:"accessory,omitempty"`
	Elements  *[]Element `json:"elements,omitempty"`
	Label     *Text      `json:"label,omitempty"`
	ImageUrl  *string    `json:"image_url,omitempty"`
	AltText   *string    `json:"alt_text,omitempty"`
}

type Element struct {
	Type                *string   `json:"type"`
	Text                *string   `json:"text,omitempty"`
	Emoji               *bool     `json:"emoji,omitempty"`
	ImageUrl            *string   `json:"image_url,omitempty"`
	AltText             *string   `json:"alt_text,omitempty"`
	Placeholder         *Text     `json:"placeholder,omitempty"`
	ActionId            *string   `json:"action_id,omitempty"`
	Options             *[]Option `json:"options,omitempty"`
	Filter              *Filter   `json:"filter,omitempty"`
	InitialConversation *string   `json:"initial_conversation,omitempty"`
	InitialUser         *string   `json:"initial_user,omitempty"`
	InitialChannel      *string   `json:"initial_channel,omitempty"`
	Value               *string   `json:"value,omitempty"`
	InitialDate         *string   `json:"initial_date,omitempty"`
	InitialTime         *string   `json:"initial_time,omitempty"`
}

type Filter struct {
	Include *[]string `json:"include,omitempty"`
}

type Accessory struct {
	Type        *string   `json:"type"`
	Placeholder *Text     `json:"placeholder,omitempty"`
	ActionId    *string   `json:"action_id,omitempty"`
	Options     *[]Option `json:"options,omitempty"`
	Value       *string   `json:"value,omitempty"`
	Url         *string   `json:"url,omitempty"`
	ImageUrl    *string   `json:"image_url,omitempty"`
	AltText     *string   `json:"alt_text,omitempty"`
	InitialDate *string   `json:"initial_date,omitempty"`
	InitialTime *string   `json:"initial_time,omitempty"`
}

type Option struct {
	Value       *string `json:"value,omitempty"`
	Text        *Text   `json:"text,omitempty"`
	Description *Text   `json:"description,omitempty"`
}
