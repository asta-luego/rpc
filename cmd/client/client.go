package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/asta-luego/rpc/pkg"
)

func main() {
	fmt.Println("Client")
	call := &definitionsrpc.Call{}

	err := rpc.Register(call)
	if err != nil {
		log.Fatal(err)
	}
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	res := 0
	args := 42
	err = client.Call("Call.FirstCall", &args, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
