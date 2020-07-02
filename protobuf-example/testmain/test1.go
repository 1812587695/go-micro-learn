package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	protoes "go-micro-learn/protobuf-example/test1"
)

func main() {

	a := &protoes.TestRequest{
		Name:   "ss",
		Age:    22,
		Height: 11,
	}
	fmt.Println(a)

	//编码
	data, _ := proto.Marshal(a)
	fmt.Println(data)

	//解码
	b := &protoes.TestRequest{}
	proto.Unmarshal(data, b)
	fmt.Println(b)
	fmt.Println(b.String())

	//获取信息
	fmt.Println(b.Name)
	fmt.Println(b.Age)
	fmt.Println(b.Height)
}
