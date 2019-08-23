package message

import "strings"

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

// SetAt 设置告警通知人，注册人的钉钉手机号
func (msg MarkdownMessage) SetAt(mobiles []string) DingMessage {
	msg.AtMobiles = mobiles
	for i := range mobiles {
		mobiles[i] = "@"+mobiles[i]
	}
	msg.Text = strings.Join(mobiles, " ") + "\n" + msg.Text
	// todo
	return msg
}


func (msg MarkdownMessage) SetAtAll(b bool) DingMessage {
	msg.At.IsAtAll = b
	return msg
}
