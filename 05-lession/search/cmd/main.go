package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"gotasks/05-lession/search/pkg/crawler"
	"gotasks/05-lession/search/pkg/crawler/spider"
	"gotasks/05-lession/search/pkg/index"
)

const fn = "search_data.txt"

func main() {

	var res []crawler.Document
	var err error

	sw := flag.String("s", "", "Search word")
	flag.Parse()

	if *sw == "" {
		fmt.Println("Search word is required")
		return
	}

	_, err = os.Stat("search_data.txt")

	if os.IsNotExist(err) {
		res, err = crawl()
	} else {
		res, err = readFile()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	makeIdx(res)
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

func writeToFile(s string) error {
	file, err := os.OpenFile(
		fn,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(s)
	return err
}

func readFile() ([]crawler.Document, error) {
	var result []crawler.Document
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(data), &result)
	return result, nil
}

func crawl() ([]crawler.Document, error) {
	urls := []string{"http://go.dev", "http:://golang.org"}
	res := []crawler.Document{}
	srv := spider.New()

	for _, u := range urls {
		docs, err := srv.Scan(u, 2)
		if err != nil {
			fmt.Println(err)
		}

		for idx, doc := range docs {
			newDocument := crawler.Document{ID: idx, URL: doc.URL, Title: doc.Title}
			res = append(res, newDocument)
		}
	}

	jStr, err := json.Marshal(res)

	if err != nil {
		return nil, err
	}

	err = writeToFile(string(jStr))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func makeIdx(docs []crawler.Document) {
	i := 0

	for _, doc := range docs {
		index.SetIndex(i, doc.Title)
		i++
	}
}
