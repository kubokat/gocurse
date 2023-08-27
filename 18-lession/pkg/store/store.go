package store

import "fmt"

type LinkStoreInterface interface {
	AddLink(short string, origin string) string
	GetLink(short string) (string, error)
}

type LinkStore struct {
	links map[string]string
}

func NewLinkStore() *LinkStore {
	return &LinkStore{
		links: make(map[string]string),
	}
}

func (ls *LinkStore) AddLink(short string, origin string) string {
	ls.links[short] = origin

	fmt.Println(ls.links)

	return short
}

func (ls *LinkStore) GetLink(short string) (string, error) {
	if _, ok := ls.links[short]; !ok {
		err := fmt.Errorf("short link %s does not exist", short)
		return "", err
	}

	return ls.links[short], nil
}
