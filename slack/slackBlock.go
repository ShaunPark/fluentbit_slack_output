package slack

type TextBlock struct {
	Type  *string `json:"type"`
	Text  *string `json:"text,omitempty"`
	Emoji *bool   `json:"emoji,omitempty"`
}

type Blocks struct {
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Type      *string      `json:"type"`
	Text      *TextBlock   `json:"text,omitempty"`
	Title     *TextBlock   `json:"title,omitempty"`
	Fields    *[]TextBlock `json:"fields,omitempty"`
	Accessory *Accessory   `json:"accessory,omitempty"`
	Elements  *[]Element   `json:"elements,omitempty"`
	Label     *TextBlock   `json:"label,omitempty"`
	ImageUrl  *string      `json:"image_url,omitempty"`
	AltText   *string      `json:"alt_text,omitempty"`
}

type Element struct {
	Type                *string    `json:"type"`
	Text                *string    `json:"text,omitempty"`
	Emoji               *bool      `json:"emoji,omitempty"`
	ImageUrl            *string    `json:"image_url,omitempty"`
	AltText             *string    `json:"alt_text,omitempty"`
	Placeholder         *TextBlock `json:"placeholder,omitempty"`
	ActionId            *string    `json:"action_id,omitempty"`
	Options             *[]Option  `json:"options,omitempty"`
	Filter              *Filter    `json:"filter,omitempty"`
	InitialConversation *string    `json:"initial_conversation,omitempty"`
	InitialUser         *string    `json:"initial_user,omitempty"`
	InitialChannel      *string    `json:"initial_channel,omitempty"`
	Value               *string    `json:"value,omitempty"`
	InitialDate         *string    `json:"initial_date,omitempty"`
	InitialTime         *string    `json:"initial_time,omitempty"`
}

type Filter struct {
	Include *[]string `json:"include,omitempty"`
}

type Accessory struct {
	Type        *string    `json:"type"`
	Placeholder *TextBlock `json:"placeholder,omitempty"`
	ActionId    *string    `json:"action_id,omitempty"`
	Options     *[]Option  `json:"options,omitempty"`
	Value       *string    `json:"value,omitempty"`
	Url         *string    `json:"url,omitempty"`
	ImageUrl    *string    `json:"image_url,omitempty"`
	AltText     *string    `json:"alt_text,omitempty"`
	InitialDate *string    `json:"initial_date,omitempty"`
	InitialTime *string    `json:"initial_time,omitempty"`
}

type Option struct {
	Value       *string    `json:"value,omitempty"`
	Text        *TextBlock `json:"text,omitempty"`
	Description *TextBlock `json:"description,omitempty"`
}
