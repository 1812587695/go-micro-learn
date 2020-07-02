package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Ta 必须首字母大小，不然rpc调用不了
type Ta int

// GetInfo 必须首字母大小，不然rpc调用不了
func (t *Ta) GetInfo(args *int, reply *int) error {

	fmt.Println(args)

	*reply = *args + 1000

	return nil
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!\n")
}

func main() {
	http.HandleFunc("/", myHandler) //	设置访问路由

	// 创建对象
	test := new(Ta)
	//rpc服务注册了一个test对象 公开方法供客户端调用
	rpc.Register(test)
	//指定rpc的传输协议 这里采用http协议作为rpc调用的载体 也可以用rpc.ServeConn处理单个连接请求
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("listen error", e)
	}
	http.Serve(l, nil)
}
