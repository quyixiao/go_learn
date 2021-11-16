package main

import (
	"fmt"
)

type SingleSender interface {
	Send(to string, msg string) error
}
type Sender interface {
	Send(to string, msg string) error
	SendAll(tos []string, msg string) error
}

type EmailSender struct {
	SmtpAddr string
}

func (e EmailSender) Send(to, msg string) error {
	fmt.Println("发送邮件给：", to, "消息内容", msg)
	return nil
}

func (e EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		e.Send(to, msg)
	}
	return nil
}

type SmsSender struct {
}

func (e SmsSender) Send(to, msg string) error {
	fmt.Println("发送邮件给：", to, "消息内容", msg)
	return nil
}

func (e SmsSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		e.Send(to, msg)
	}
	return nil
}

func do(sender Sender) {
	sender.Send("领导", "工作报告")
}

type WeChatSender struct {
	ID int
}

func (e WeChatSender) Send(to, msg string) error {
	fmt.Println("发送微信给：", to, "消息内容", msg)
	return nil
}

func (e *WeChatSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		e.Send(to, msg)
	}
	return nil
}

func main() {
	var sender Sender = EmailSender{}
	fmt.Printf("%T ,%#v \n", sender, sender)
	sender.Send("张三", "hello word!")
	sender.SendAll([]string{"张三", "王五"}, "hello word!")

	var smsSender Sender = SmsSender{}
	fmt.Printf("%T ,%#v \n", smsSender, smsSender)
	sender.Send("张三", "hello word!")
	sender.SendAll([]string{"张三", "王五"}, "hello word!")

	do(sender)
	do(smsSender)

	pointerSender := EmailSender{}
	do(&pointerSender)

	pointSmsSender := SmsSender{}
	do(&pointSmsSender)

	weChatSender := &WeChatSender{}

	weChatSender.Send("11111111", "2222") //发送微信给： 11111111 消息内容 2222

	//发送微信给： 11111 消息内容 3333333
	//发送微信给： 22222 消息内容 3333333
	weChatSender.SendAll([]string{"11111", "22222"}, "3333333")

	var ssender SingleSender = weChatSender
	ssender.Send("哈哈", "你是") //发送微信给： 哈哈 消息内容 你是

	//强制类型转换
	sender01 := ssender.(*WeChatSender)
	fmt.Printf("%T ,%v \n", sender01, sender01) //*main.WeChatSender ,&{}
	sender01.Send("张三", "hello")                //发送微信给： 张三 消息内容 hello

	//发送微信给： 小凡 消息内容 hello world
	//发送微信给： 李四 消息内容 hello world
	sender01.SendAll([]string{"小凡", "李四"}, "hello world")
	fmt.Println(sender01.ID)

	esender01 := ssender.(*WeChatSender)
	fmt.Printf("%T ,%T \n", esender01, esender01) //*main.WeChatSender ,*main.WeChatSender
	//根据具体的类型去做处理,只能用在switch的子语句中
	switch sender.(type) {
	case EmailSender:
		fmt.Println("EmailSender")
	}

}
