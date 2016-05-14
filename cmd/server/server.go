package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

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
	ln, err := net.Listen("tcp", ":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	err = http.Serve(ln, nil)
	if err != nil {
		log.Fatal(err)
	}

}
