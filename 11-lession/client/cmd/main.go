package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	client()
}

func client() {
	conn, err := net.Dial("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	r := bufio.NewScanner(os.Stdin)

	for {
		r.Scan()
		req := r.Text()

		_, err := conn.Write([]byte(req + "\n"))
		if err != nil {
			continue
		}

		res, _, err := bufio.NewReader(conn).ReadLine()
		if err != nil {
			continue
		}

		fmt.Println("Response:", string(res))
	}

}
