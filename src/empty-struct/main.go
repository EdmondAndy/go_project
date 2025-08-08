package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Codec interface {
    Encode(w io.Writer, v interface{}) error
    Decode(r io.Reader, v interface{}) error
}    

type jsonCodec struct{}
func (jsonCodec) Encode(w io.Writer, v interface{}) error {
    return json.NewEncoder(w).Encode(v)
}
func (jsonCodec) Decode(r io.Reader, v interface{}) error {
    return json.NewDecoder(r).Decode(v)
}    

var JSON Codec = jsonCodec{}

func main() {
    sobj := struct {
        S1 string `json:"s1"`
        K3 string `json:"k3"`
    }{}
    ss := `{"s1": "v1", "k3": "vv3"}`
    err := JSON.Decode(strings.NewReader(ss), &sobj)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(sobj)
}