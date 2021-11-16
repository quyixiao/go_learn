package main

import "fmt"

type Host struct {
	Id   string
	Name string
}

type Clound interface {
	GetList() []Host
	Start(Id string) error
	Stop(Id string) error
}

type Tenant struct {
}

func (tenant Tenant) GetList() []Host {
	return []Host{}
}

func (tenant Tenant) Start(Id string) error {
	fmt.Println("腾讯start ")
	return nil
}

func (tenant Tenant) Stop(Id string) error {
	fmt.Println("腾讯stop")
	return nil
}

type Aliyun struct {
}

func (tenant Aliyun) GetList() []Host {
	return []Host{}
}

func (tenant Aliyun) Start(Id string) error {
	fmt.Println("阿里云start ")
	return nil
}

func (tenant Aliyun) Stop(Id string) error {
	fmt.Println("阿里云stop")
	return nil
}

func main() {
	var clound Clound
	clound = Tenant{}

	clound.Start("89329")
}
