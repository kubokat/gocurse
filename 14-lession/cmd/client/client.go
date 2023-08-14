package main

import (
	"fmt"
	"log"
	"net/rpc"

	msg "gotasks/14-lession/pkg/message"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8090")
	if err != nil {
		log.Fatal(err)
	}

	var res string
	client.Call("Server.Send", []string{"first", "second", "last"}, &res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	var messages []msg.Message
	client.Call("Server.Messages", struct{}{}, &messages)

	fmt.Printf("%+v\n", messages)
}
