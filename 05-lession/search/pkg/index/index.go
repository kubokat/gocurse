package index

import (
	"strings"
)

var Idx = make(map[string][]int)
var wordsArr []string

func SetIndex(index int, words string) {
	wordsArr = strings.Split(words, " ")
	for _, k := range wordsArr {
		Idx[k] = append(Idx[k], index)
	}
}
