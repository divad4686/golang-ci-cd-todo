package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"
)

type todoItemSend struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Text      string `json:"text"`
}

type todoItemResponse struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	Completed bool   `json:"completed"`
	Text      string `json:"text"`
}

func TestSum(t *testing.T) {
	total := 5 + 5
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestApi(t *testing.T) {

	item := todoItemSend{
		Title:     "List",
		Completed: false,
		Text:      "hacer cosas",
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = "http://localhost:8080"
	}

	jsonitem, err := json.Marshal(item)
	if err != nil {
		t.Error("Could not create json")
	}

	posturl := host + "/todos/"
	resp, err := http.Post(posturl, "application/json; charset=utf-8", bytes.NewBuffer(jsonitem))

	if err != nil {
		t.Error(err.Error())
	}

	defer resp.Body.Close()

	result := todoItemResponse{}
	json.NewDecoder(resp.Body).Decode(&result)

	if result.URL == "" {
		t.Error("Imcomplete response")
	}

	resp, err = http.Get(result.URL)
	if err != nil {
		t.Error(err.Error())
	}

	getresult := todoItemResponse{}
	json.NewDecoder(resp.Body).Decode(&getresult)

	if getresult.Title != result.Title || getresult.Completed != result.Completed {
		t.Error("Different result")
	}
}
