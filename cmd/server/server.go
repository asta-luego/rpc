package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/asta-luego/rpc/pkg"
)

func main() {
	fmt.Println("Serveur")
	call := &definitionsrpc.Call{}

	err := rpc.Register(call)
	if err != nil {
		log.Fatal(err)
	}

	rpc.HandleHTTP()
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	err = http.Serve(ln, nil)
	if err != nil {
		log.Fatal(err)
	}

}
