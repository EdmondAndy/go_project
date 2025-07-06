package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockClient struct {
	GetResponseOutput *http.Response
	PostResponseOutput *http.Response
}

func (m MockClient) Get(url string) (resp *http.Response, err error) {
	// This is a mock implementation of the Get method.
	// In a real test, you would return a mock response or an error.
	return m.GetResponseOutput, nil
}

func (m MockClient) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return m.PostResponseOutput, nil
}

func TestDoGetRequest(t *testing.T) {
	// This function is a placeholder for the actual test implementation.
	// It should contain the logic to test the DoGetRequest method of the API interface.
	// You can use a mock server or a real server for testing purposes.
	words := WordsPage{
		Page: Page{"words"},
		Words: Words{
			Input: "abc",
			Words: []string{"a", "b"},
		},
	}
	wordsBytes, err := json.Marshal(words)
	if err != nil {
		t.Fatalf("Failed to marshal words: %s", err)
	}

	apiInstance := API{
		Options: Options{},
		Client: MockClient{
			GetResponseOutput: &http.Response{
				StatusCode: 200,
				Body: io.NopCloser(bytes.NewReader(wordsBytes)),
			},
		},
	}
	response, err := apiInstance.DoGetRequest("http://localhost/words")
	if err != nil {
		t.Errorf("DoGetRequest error: %s", err)
	}
	if response == nil {
		t.Fatalf("response is empty")
	}
	if response.GetResponse() != strings.Join([]string{"Words: a", "b"}, ", "){
		t.Errorf("Unexpected response: %s", response.GetResponse())
	}
}