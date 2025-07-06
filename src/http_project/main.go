package main

import (
	"fmt"
	"io"
	"log"
)

// Words struct to hold the JSON response
type MySlowReader struct {
	Contents string
	pos int
}

func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.pos+1 <= len(m.Contents){
		n := copy(p, m.Contents[m.pos:m.pos+1])
		m.pos ++
		return n, nil

	}
	return 0, io.EOF // Simulate end of file immediately
}

func main() {

	MySlowReaderInstance := &MySlowReader{
		Contents: "hello world",
	}

	out, error := io.ReadAll(MySlowReaderInstance)
	
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("output: %s\n", out)

	// error = json.Unmarshal(body, &words)
	// if error != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("JSON Parsed\nPage: %s\nWords: %v\n", words.Page, strings.Join(words.Words, ", "))
}