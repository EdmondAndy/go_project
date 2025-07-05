package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Words struct to hold the JSON response
type Words struct {
	Page string `json:"page"`
	Input string `json:"input"`
	Words []string `json:"words"`
}

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}
	if myUrl, error := url.ParseRequestURI(args[1]); error != nil {
		fmt.Printf("Invalid URL: %s\n, error is %s\n", args[1], error)
		os.Exit(1)
	} else {
		fmt.Printf("Valid URL: %s\n", myUrl.String())
	}
	if response, error := http.Get(args[1]); error != nil {
		fmt.Printf("Error fetching URL: %s\n, error is %s\n", args[1], error)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		body, error := io.ReadAll(response.Body)
		if error != nil {
			fmt.Printf("Error reading response body: %s\n, error is %s\n", args[1], error)
			os.Exit(1)
		}
		fmt.Printf("Response Code: %d\n", response.StatusCode)
		fmt.Printf("Response Body: %s\n", string(body))
	}

	// error = json.Unmarshal(body, &words)
	// if error != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("JSON Parsed\nPage: %s\nWords: %v\n", words.Page, strings.Join(words.Words, ", "))
}