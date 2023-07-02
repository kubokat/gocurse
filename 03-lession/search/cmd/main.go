package main

import (
	"flag"
	"fmt"
	"sort"

	"gotasks/03-lession/search/pkg/crawler"
	"gotasks/03-lession/search/pkg/crawler/spider"
	"gotasks/03-lession/search/pkg/index"
)

func main() {

	sw := flag.String("s", "", "Search word")
	flag.Parse()

	if *sw == "" {
		fmt.Println("Search word is required")
		return
	}

	i := 0

	urls := []string{"http://go.dev", "http:://golang.org"}
	res := []crawler.Document{}
	srv := spider.New()

	for _, u := range urls {
		docs, err := srv.Scan(u, 2)
		if err != nil {
			fmt.Println(err)
		}

		for _, doc := range docs {
			newDocument := crawler.Document{ID: i, URL: doc.URL, Title: doc.Title}
			index.SetIndex(i, doc.Title)
			res = append(res, newDocument)
			i++
		}
	}

	sIds := index.Idx[*sw]

	for _, id := range sIds {
		if id < len(res) {
			dIdx := sort.Search(len(res), func(i int) bool {
				return res[i].ID >= id
			})

			fmt.Println(res[dIdx])
		}
	}
}
