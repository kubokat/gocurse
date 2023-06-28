package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"gotasks/02-lession/search/pkg/crawler/spider"
)

func main() {

	sw := flag.String("s", "", "Search word")
	flag.Parse()

	if *sw == "" {
		log.Fatal("Search word is required")
	}

	urls := []string{"http://go.dev", "http:://golang.org"}
	srv := spider.New()

	for _, u := range urls {
		docs, err := srv.Scan(u, 2)
		if err != nil {
			log.Fatal(err)
		}

		for _, doc := range docs {
			if strings.Contains(doc.Title, *sw) {
				fmt.Println(doc.Title)
			}
		}
	}
}
