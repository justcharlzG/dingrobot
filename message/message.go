// https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq
package message

import (
	"encoding/json"
	"strings"
)

type DingType string

const (
	MsgText       DingType = "text"
	MsgLink       DingType = "link"
	MsgMarkdown   DingType = "markdown"
	MsgActionCard DingType = "actionCard"
	MsgFeedCard   DingType = "feedCard"
)

// DingMessage a interface robot message
type DingMessage interface {
	MessageType() DingType
}

// AtMobiles a interface set at , the `@` feature support text and markdown
type AtMobiles interface {
	// set mobiles
	SetAt([]string) DingMessage
	SetAtAll() DingMessage
}

// Message send body
type Message struct {
	MsgType DingType `json:"msgtype"`
	DingMessage
}

// MarshalJSON 实现数据json转换，使用了interface [DingMessage], 技术水平限制，先用switch转换
func (m Message) MarshalJSON() ([]byte, error) {
	switch m.MsgType {
	case MsgText:
		return json.Marshal(struct {
			MsgType DingType `json:"msgtype"`
			TextMessage
		}{
			m.MsgType,
			m.DingMessage.(TextMessage),
		})
	case MsgLink:
		return json.Marshal(struct {
			MsgType DingType `json:"msgtype"`
			LinkMessage
		}{
			m.MsgType,
			m.DingMessage.(LinkMessage),
		})
	case MsgMarkdown:
		return json.Marshal(struct {
			MsgType DingType `json:"msgtype"`
			MarkdownMessage
		}{
			m.MsgType,
			m.DingMessage.(MarkdownMessage),
		})
	case MsgActionCard:
		return json.Marshal(struct {
			MsgType DingType `json:"msgtype"`
			ActionCardMessage
		}{
			m.MsgType,
			m.DingMessage.(ActionCardMessage),
		})
	default:
		return nil, nil
	}
}

// Response robot response
type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// TextMessage text message
type TextMessage struct {
	TextContent `json:"text"`
	At          `json:"at,omitempty"`
}

func (msg TextMessage) SetAtAll(b bool) DingMessage {
	msg.At.IsAtAll = b
	return msg
}

// SetAt 设置告警通知人，注册人的钉钉手机号
func (msg TextMessage) SetAt(mobiles []string) DingMessage {
	msg.AtMobiles = mobiles
	for i := range mobiles {
		mobiles[i] = "@"+mobiles[i]
	}
	msg.Content = strings.Join(mobiles, " ") + "\n" + msg.Content
	// todo
	return msg
}


func (TextMessage) MessageType() DingType {
	return MsgText
}

type TextContent struct {
	Content string `json:"content"`
}

type At struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

// LinkMessage link message
type LinkMessage struct {
	LinkContent `json:"link"`
}

func (LinkMessage) MessageType() DingType {
	return MsgLink
}

type LinkContent struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageURL string `json:"messageUrl"`
	PicURL     string `json:"picUrl,omitempty"`
}



// ActionCardMessage action
type ActionCardMessage struct {
	ActionCardContent `json:"actionCard"`
}

func (ActionCardMessage) MessageType() DingType {
	return MsgActionCard
}

type ActionCardContent struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
	BtnOrientation string `json:"btnOrientation,omitempty"`
	HideAvatar     string `json:"hideAvatar,omitempty"`
}

