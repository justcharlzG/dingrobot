package message

// MarkdownMessage mark down message
type MarkdownMessage struct {
	MarkdownContent `json:"markdown"`
	At              `json:"at,omitempty"`
}

func (MarkdownMessage) MessageType() DingType {
	return MsgMarkdown
}

type MarkdownContent struct {
	Title string       `json:"title"`
	Text  string `json:"text"`
}
