package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// The client should be able to perform a request
type Client interface {
	PerformRequest(r *http.Request) (*http.Response, error)
}

// Creates a HTTP client.
func HTTPClient(secure bool, timeout int) *http.Client {
	return &http.Client{
		Timeout:   time.Duration(timeout * int(time.Second)),
		Transport: tlsConfigTransport(secure),
	}
}

// Returns transport based on TLS configuration. Client verifies the server's certificate and host name based on secure.
func tlsConfigTransport(secure bool) *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: secure,
		},
	}
}

// Extract base URL and return a net/url URL.
func extractBaseURL(baseURL string) (*url.URL, error) {
	// Raw URL to URL struct.
	url, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}

	// Reject any other protocol except HTTP & HTTPS.
	if url.Scheme != "http" && url.Scheme != "https" {
		return nil, errors.New("Invalid protocol.")
	}

	return url, nil
}

// Parse config file and extract tests.
// TODO: Implement proposed Tests struct
func extractTests(file string) (Tests, error) {
	content, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	var tests Tests

	// Unmarshal and store to tests
	err = json.Unmarshal(content, &tests)

	if err != nil {
		return nil, err
	}

	// Handle invalid tests
	for _, test := range tests {
		if test.URL == "" || test.StatusCode == 0 {
			return nil, errors.New("URL not found. Please verify your test.")
		}
	}

	return tests, nil
}

// TODO: This is a placeholder
// There needs to be a lot of work to be done here.
func prettifyResult(result Result) string {
	output := fmt.Sprintf("Expected Status Code: %d\nStatus Code: %d\nURL:%s", result.StatusCodeExpected, result.StatusCode, result.URL)

	return output
}
