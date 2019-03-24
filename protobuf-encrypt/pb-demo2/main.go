package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"pb-demo2/contact"
)

func write() {
	p1 := &contact.Person{
		Id:   1,
		Name: "小张",
		Phones: []*contact.Phone{
			{Type: contact.PhoneType_HOME, Number: "11111111"},
		},
	}
	p2 := &contact.Person{
		Id:   2,
		Name: "小王",
		Phones: []*contact.Phone{
			{Type: contact.PhoneType_HOME, Number: "333333333"},
			{Type: contact.PhoneType_WORK, Number: "444444444"},
		},
	}

	//创建地址簿
	book := &contact.ContactBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)

	//编码数据
	data, _ := proto.Marshal(book)
	//把数据写入文件
	ioutil.WriteFile("./contact.txt", data, os.ModePerm)
}

func read() {
	//读取文件数据
	data, _ := ioutil.ReadFile("./contact.txt")
	book := &contact.ContactBook{}
	//解码数据
	proto.Unmarshal(data, book)
	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)
		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
	}
}

func main() {
	write()
	read()
}
