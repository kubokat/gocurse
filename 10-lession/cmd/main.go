package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	name   string
	points int
	ch     chan string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	ch := make(chan string)
	p := Player{name: "Player1", points: 0, ch: ch}
	p2 := Player{name: "Player2", points: 0, ch: ch}

	wg.Add(2)
	go p.pass(&wg)
	go p2.pass(&wg)

	ch <- "begin"
	wg.Wait()

	fmt.Printf("%s points %d vs %s points %d", p.name, p.points, p2.name, p2.points)
}

func (p *Player) pass(wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range p.ch {
		if msg == "stop" {
			fmt.Println(msg)
			p.ch <- "begin"
			continue
		} else if msg == "begin" {
			fmt.Println(msg)
			p.ch <- "Ping"
			continue
		} else {
			fortune := rand.ExpFloat64() < 0.2

			if !fortune {
				fmt.Println(p.name, msg)

				if msg == "Pong" {
					p.ch <- "Ping"
				} else {
					p.ch <- "Pong"
				}
			} else {
				if p.points == 10 {
					close(p.ch)
					break
				}

				p.points++
				fmt.Printf("%s win\n", p.name)
				p.ch <- "stop"
			}
		}
	}
}
