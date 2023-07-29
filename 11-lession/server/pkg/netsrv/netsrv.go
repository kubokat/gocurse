package netsrv

import (
	"bufio"
	"encoding/json"
	"gotasks/11-lession/server/pkg/cache"
	"gotasks/11-lession/server/pkg/crawler"
	"gotasks/11-lession/server/pkg/crawler/spider"
	"gotasks/11-lession/server/pkg/index"
	"log"
	"net"
	"os"
	"sort"
	"sync"
)

func handler(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)

	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		f := search()
		sIds := index.Idx[string(msg)]

		for _, id := range sIds {
			if id < len(f) {
				dIdx := sort.Search(len(f), func(i int) bool {
					return f[i].ID >= id
				})

				_, err = conn.Write([]byte(f[dIdx].Title + "\n"))
				if err != nil {
					return
				}
			}
		}
	}
}

func Listen() {
	listener, err := net.Listen("tcp", "0.0.0.0:8000")

	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handler(conn)
	}
}

func search() []crawler.Document {
	var res []crawler.Document
	var err error

	_, err = os.Stat(cache.Fn())

	if os.IsNotExist(err) {
		res, err = crawl()
	} else {
		res, err = cache.Read()
	}

	if err != nil {
		log.Fatal(err)
	}

	index.MakeIdx(res)
	return res
}

func crawl() ([]crawler.Document, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	urls := []string{"http://go.dev", "http:://golang.org"}
	spd := spider.New()
	res := []crawler.Document{}

	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			docs, err := spd.Scan(url, 1)

			if err != nil {
				log.Fatal(err)
			}

			for idx, doc := range docs {
				mu.Lock()
				res = append(res, crawler.Document{ID: idx, URL: doc.URL, Title: doc.Title})
				mu.Unlock()
			}
		}(u)
	}

	wg.Wait()

	jStr, err := json.Marshal(res)

	if err != nil {
		return nil, err
	}

	err = cache.Write(string(jStr))

	if err != nil {
		return nil, err
	}

	return res, nil
}
