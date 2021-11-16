package main

import "fmt"

type Sender interface {
	Sender(msg string) error
}

type Reciver interface {
	Recive() (string, error)
}

type Client interface {
	Sender
	Reciver

	Open() error
	Close() error
}

type MSNClient struct {
}


func (c MSNClient) Open() error {
	fmt.Println("open： " )
	return nil
}

func (c MSNClient) Close() error {
	fmt.Println("close " )
	return nil
}
func (c MSNClient) Sender(msg string) error {
	fmt.Println("sender： " + msg)
	return nil
}

func (c MSNClient) Recive() (string, error) {
	fmt.Println("reciver ")
	return "", nil
}

func main() {

	msn := MSNClient{}

	var s Sender = msn
	var r Reciver = msn
	var c Client = msn

	s.Sender("哈哈")
	r.Recive()
	c.Sender("你好")
	c.Recive()




	defer msn.Close()




}
