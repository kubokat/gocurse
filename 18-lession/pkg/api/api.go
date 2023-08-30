package api

import (
	"encoding/json"
	"gotasks/18-lession/pkg/shorter"
	"gotasks/18-lession/pkg/store"

	"net/http"
)

type API struct {
	LinkStore store.LinkStoreInterface
}

type RequestBody struct {
	Url string `json:"url"`
}

func (api *API) MakeUrl(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response("shorter", api.LinkStore.AddLink(shorter.CreateSLink(requestBody.Url), requestBody.Url), w)
}

func (api *API) GetOrigin(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	link, err := api.LinkStore.GetLink(requestBody.Url)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response("original", link, w)
}

func response(key string, link string, w http.ResponseWriter) {
	response := map[string]string{
		key: link,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
