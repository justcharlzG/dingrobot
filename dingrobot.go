package dingrobot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/guowenshuai/dingrobot/message"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Roboter interface {
	Send(message.DingMessage) error
}

type Robot struct {
	Webhook string
}

// NewRobot returns a dingRoboter client
func NewRobot(webhook string) Roboter {
	return &Robot{Webhook: webhook}
}

// Send send a message which must DingMessage
func (r *Robot) Send(msg message.DingMessage) error {
	// generate a request id
	logrus.SetLevel(logrus.DebugLevel)
	rid := strconv.Itoa(rand.Int())

	uuid, err := uuid.NewV4()
	if err != nil {
		logrus.Debugf("[%s] generate uuid error: %s", rid, err.Error())
	}
	rid = uuid.String()[:6]


	logrus.Infof("[%s] send [%s] Type message to ding", rid, msg.MessageType())


	payload := message.Message{
		MsgType: msg.MessageType(),
		DingMessage: msg,
	}

	m, err := json.Marshal(payload)
	if err != nil {
		logrus.Debugf("[%s] json marshal error: %s", rid, err.Error())
		return errors.Wrap(err, "json marshal message error ")
	}

	// 发送数据
	resp, err := http.Post(r.Webhook, "application/json", bytes.NewReader(m))
	if err != nil {
		logrus.Debugf("[%s] post to ding error: %s", rid, err.Error())
		return errors.Wrap(err, "post to ding error")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Debugf("[%s] read response error: %s", rid, err.Error())
		return errors.Wrap(err, "read response error")
	}

	var rBody message.Response
	err = json.Unmarshal(data, &rBody)
	if err != nil {
		logrus.Debugf("[%s] json unmarshal response error: %s", rid, err.Error())
		return errors.Wrap(err, "json unmarshal response error ")
	}
	if rBody.ErrCode != 0 {
		return fmt.Errorf("dingrobot send failed: code [%d], msg [%v]", rBody.ErrCode, rBody.ErrMsg)
	}

	return nil
}
