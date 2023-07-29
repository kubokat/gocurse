package cache

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"gotasks/11-lession/server/pkg/crawler"
)

func Fn() string {
	return "search_data.txt"
}

func Write(s string) error {
	file, err := os.OpenFile(
		Fn(),
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

func Read() ([]crawler.Document, error) {
	var result []crawler.Document
	file, err := os.Open(Fn())
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
