package dingrobot

import (
	"testing"

	"github.com/guowenshuai/dingrobot/message"
	md "github.com/guowenshuai/simpleMarkDown"
)

func TestRobot_Send(t *testing.T) {
	type fields struct {
		Webhook string
	}
	builder := md.NewBuilder()

	mb := []string{"18810975701", "17703063443"}
	type args struct {
		msg message.DingMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   string(message.MsgText),
			fields: fields{"https://oapi.dingtalk.com/robot/send?access_token=93839c32719109fb1eac0b2118e471addd964379721941985c3abde6e2607190"},
			args: args{
				msg: message.TextMessage{
					TextContent: message.TextContent{"this is text messsage"},
				}.SetAtAll(true),
			},
			wantErr: false,
		},
		{
			name:   string(message.MsgLink),
			fields: fields{"https://oapi.dingtalk.com/robot/send?access_token=93839c32719109fb1eac0b2118e471addd964379721941985c3abde6e2607190"},
			args: args{
				msg: message.LinkMessage{
					LinkContent: message.LinkContent{
						Text:       "中国和谐号动车组",
						Title:      "时代的火车向前开",
						PicURL:     "http://a0.att.hudong.com/85/17/01300000241358122593175732033.jpg",
						MessageURL: "https://baike.sogou.com/v267779.htm?fromTitle=%E5%92%8C%E8%B0%90%E5%8F%B7%E5%8A%A8%E8%BD%A6%E7%BB%84",
					},
				},
			},
			wantErr: false,
		},
		{
			name:   string(message.MsgMarkdown),
			fields: fields{"https://oapi.dingtalk.com/robot/send?access_token=93839c32719109fb1eac0b2118e471addd964379721941985c3abde6e2607190"},
			args: args{
				msg: message.MarkdownMessage{
					MarkdownContent: message.MarkdownContent{
						Title: "时代的火车向前开",
						Text: builder.AddHeader(md.Level2, "北京天气").
							AddRow(md.NewRow("9度,").Bold("西北风1级,").Italic("空气良89，相对温度73%")).
							AddRow(md.NewRow("").Picture("天气", "https://cms-bucket.nosdn.127.net/catchpic/a/ac/ac5ed48a4e9a9b153216af8a54c42c6e.jpg?imageView&thumbnail=550x0")).
							String(),
					},
				}.SetAt(mb),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				Webhook: tt.fields.Webhook,
			}
			if err := r.Send(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Robot.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
