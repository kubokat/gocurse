package webapp

import (
	"encoding/json"
	"gotasks/11-lession/server/pkg/crawler"
	"gotasks/11-lession/server/pkg/index"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type MockCache struct {
	MockRead  func() ([]crawler.Document, error)
	MockWrite func(s string) error
}

func (mc *MockCache) Read() ([]crawler.Document, error) {
	return mc.MockRead()
}

func (mc *MockCache) Write(s string) error {
	return mc.MockWrite(s)
}

func TestViewDocs(t *testing.T) {
	mockData := []crawler.Document{{ID: 1, URL: "http://example.com", Title: "Example Title"}}
	cache := &MockCache{
		MockRead: func() ([]crawler.Document, error) {
			return mockData, nil
		},
	}

	req, err := http.NewRequest("GET", "/docs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		viewDocs(w, r, cache)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var gotData map[string]interface{}
	var expectedData map[string]interface{}

	json.Unmarshal([]byte(rr.Body.String()), &gotData)
	expectedStr := `[{"ID":1,"URL":"http://example.com","Title":"Example Title"}]`
	json.Unmarshal([]byte(expectedStr), &expectedData)

	if !reflect.DeepEqual(gotData, expectedData) {
		t.Errorf("handler returned unexpected body: got %v want %v", gotData, expectedData)
	}
}

func TestViewIndex(t *testing.T) {
	index.Idx = map[string][]int{
		"example": {1, 2, 3},
	}

	req, err := http.NewRequest("GET", "/index", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(viewIndex)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedStr := `{"example":[1,2,3]}`
	if rr.Body.String() != expectedStr {
		t.Errorf("handler returned unexpected body: got %#v want %#v", rr.Body.String(), expectedStr)
	}
}
