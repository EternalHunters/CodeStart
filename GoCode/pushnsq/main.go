// pushnsq project main.go
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
)

type NSQSender interface {
	Send(topic string, data []byte) error
}

type sender struct {
	host string `NSQ服务器地址`
}

func NewSender(host string) NSQSender {
	return &sender{host}
}

//生成发送消息URL
func (i *sender) getUrl(topic string) string {
	return fmt.Sprintf("http://%s/pub?topic=%s", i.host, topic)
}

//发送消息
func (i *sender) Send(topic string, data []byte) error {
	url := i.getUrl(topic)
	response, err := http.Post(url, "application/gpb", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Send NSQ Message failed ! url: %s err: %v", url, err)
		return err
	}
	result, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		fmt.Printf("Send NSQ Message Read Response Body failed! url: %s err: %v", url, err)
		return err
	}
	if string(result) != "OK" {
		return errors.New("Send NSQ Message NSQ Response Error, result:" + string(result))
	}
	return nil
}

func main() {
	host := "127.0.0.1:4151"
	topic := "ocs_msgpush_custom"
	//topic := "ocs_msgpush_sms"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	msgId := r.Intn(10000000)

	//body := NewSmsBody()
	//body := NewAppBody()
	//body := NewEnterpriseAppBody()
	body := NewCustomBody()
	//body := NewAppFaultBody()
	msgBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("json msg body fail, ", err)
		return
	}
	message := Message{
		Header: &Header{
			Id:   proto.Int32(int32(msgId)),
			Type: proto.Int32(1),
			//			ServiceName: proto.String("/device/fault"),
			//			From:        proto.String("sys_ocs_recr"),
			ServiceName: proto.String("/msgpusher/push-task"),
			From:        proto.String("ocs_msgpush_pusher"),
			FromType:    proto.Int32(2),
			To:          proto.String("notice"),
			ToType:      proto.Int32(3),
			SessionId:   proto.String(fmt.Sprintf("sessionid%d", msgId)),
			CreateTime:  proto.Int64(time.Now().Unix()),
			Version:     proto.String("1.00.00"),
		},
		Body: msgBody,
	}
	messageJson, err := proto.Marshal(&message)
	if err != nil {
		fmt.Println("json message fail, ", err)
	}

	sender := NewSender(host)
	resErr := sender.Send(topic, messageJson)
	fmt.Println("=====send ok====")
	fmt.Println(resErr)
}
