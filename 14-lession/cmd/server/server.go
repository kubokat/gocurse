package main

import (
	msg "gotasks/14-lession/pkg/message"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Server struct {
	messages []msg.Message
}

func (s *Server) Send(messages []string, res *string) error {
	for _, m := range messages {
		s.messages = append(s.messages, msg.Message{Cont: m, Time: time.Now()})
	}

	*res = "Messages added"

	return nil
}

func (s *Server) Messages(req struct{}, res *[]msg.Message) error {
	*res = s.messages

	return nil
}

func main() {
	srv := new(Server)
	err := rpc.Register(srv)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go rpc.ServeConn(conn)
	}
}
