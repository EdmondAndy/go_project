package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

type MockRoundTripper struct {
	RoundTripperOutput *http.Response
}

func (m MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Header.Get("Authorization") != "Bearer abc" {
		return nil, fmt.Errorf("wrong authorization header: %s", req.Header.Get("Authorization"))
	}
	return m.RoundTripperOutput, nil
}

func TestRoundTrip(t *testing.T) {
	// This function is a placeholder for the actual test implementation.
	// It should contain the logic to test the RoundTrip method of the MyJWTTransport struct.
	// You can use a mock server or a real server for testing purposes.
	// The test should verify that the Authorization header is set correctly with the token.
	loginResponse := LoginResponse{
		Token: "abc",
	}
	loginResponseBytes, err := json.Marshal(loginResponse)
	if err != nil {
		t.Errorf("marshal error: %s", err)
	}
	MyJWTTransport := MyJWTTransport{
		transport: MockRoundTripper{
			RoundTripperOutput: &http.Response{
				StatusCode: http.StatusOK,
			},
		},
		HTTPClient: MockClient{
			PostResponseOutput: &http.Response{
				StatusCode: 200,
				Body: 	 io.NopCloser(bytes.NewReader(loginResponseBytes)),
			},
		},
		password: "xyz",
	}
	req := &http.Request{
		Header: make(http.Header),
	}
	res, err := MyJWTTransport.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip error: %s", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}
}