// test_gomail project main.go
package main

import (
	"fmt"

	"sp1719.com/bastet/core/mail"
)

func main() {
	sender := mail.NewMailSender("smtpdm.aliyun.com", 25, "noreply@customer.1719.cn", "*****")
	e := &mail.Email{From: "noreply@customer.1719.cn", FromName: "xioass", Subject: "hello", Body: "Hello"}
	addresses := make([]mail.EmailUser, 0, 1)
	addresses = append(addresses, mail.EmailUser{"xiaoss@wondershare.cn", "xiao"})

	e.Addresses = addresses

	if err := sender.Send(e); err != nil { //发送邮件失败，返回重试错误进行重发
		fmt.Println("SendEmailTaskCtrl send mail failed! err", err.Error())
	}
}
