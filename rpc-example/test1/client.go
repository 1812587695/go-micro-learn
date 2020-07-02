package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var t int
	err = client.Call("Ta.GetInfo", 1000, &t)

	if err != nil {
		log.Fatal("client call:", err)
	}

	fmt.Println("rpc:", t)
}
