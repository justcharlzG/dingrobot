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
	Title string `json:"title"`
	Text  string `json:"text"`
}

// SetAt 设置告警通知人，注册人的钉钉手机号
func (msg MarkdownMessage) SetAt(mobiles []string) DingMessage {
	msg.AtMobiles = mobiles
	text := ""
	for i := range mobiles {
		text += "@" + mobiles[i] + " "
	}
	msg.Text = text + "\n" + msg.Text
	return msg
}

func (msg MarkdownMessage) SetAtAll(b bool) DingMessage {
	msg.At.IsAtAll = b
	return msg
}
