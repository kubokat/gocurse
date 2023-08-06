package webapp

import (
	"encoding/json"
	"gotasks/11-lession/server/pkg/cache"
	"gotasks/11-lession/server/pkg/index"
	"net/http"
)

func Start(port string, c cache.Cache) {
	r := http.NewServeMux()
	r.HandleFunc("/index", viewIndex)
	r.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		viewDocs(w, r, c)
	})
	http.ListenAndServe(port, r)
}

func viewIndex(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(index.Idx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func viewDocs(w http.ResponseWriter, r *http.Request, c cache.Cache) {
	res, err := c.Read()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
