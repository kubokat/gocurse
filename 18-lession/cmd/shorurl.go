package main

import (
	"gotasks/18-lession/pkg/api"
	"gotasks/18-lession/pkg/store"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	linkStore := store.NewLinkStore()

	api := &api.API{
		LinkStore: linkStore,
	}

	r.HandleFunc("/make", api.MakeUrl).Methods("POST")
	r.HandleFunc("/origin", api.GetOrigin).Methods("POST")

	http.ListenAndServe(":8121", r)
}
