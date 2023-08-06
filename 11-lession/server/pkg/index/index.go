package index

import (
	"gotasks/11-lession/server/pkg/crawler"
	"strings"
)

var Idx = make(map[string][]int)
var wordsArr []string

func setIndex(index int, words string) {
	wordsArr = strings.Split(words, " ")
	for _, k := range wordsArr {
		Idx[k] = append(Idx[k], index)
	}
}

func MakeIdx(docs []crawler.Document) {
	i := 0

	for _, doc := range docs {
		setIndex(i, doc.Title)
		i++
	}
}
