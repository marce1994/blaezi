// Copyright 2021 Aadhav Vignesh

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	// Do sends an HTTP request and returns an HTTP response. Adapted for being compliant withnet/http.
	Do(r *http.Request) (*http.Response, error)
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
		return nil, errors.New("invalid protocol used in base URL")
	}

	return url, nil
}

// Parse config file and extract tests.
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
		if test.URL == "" || test.StatusCodeExpected == 0 {
			return nil, errors.New("status code/URL not found. Please verify your test")
		}
	}

	return tests, nil
}

// TODO: This is a placeholder
// There needs to be a lot of work to be done here.
func prettifyResult(result Result) string {
	output := fmt.Sprintf("Expected Status Code: %d\nStatus Code: %d\nURL: %s", result.Test.StatusCodeExpected, result.StatusCode, result.Test.URL)

	// if content is not empty, return result with content
	if result.Test.Content != "" {
		output = fmt.Sprintf("Expected Status Code: %d\nStatus Code: %d\nURL: %s\nContent: %s", result.Test.StatusCodeExpected, result.StatusCode, result.Test.URL, result.Test.Content)
	}

	// prettify using color
	color := 32
	status := "PASS"

	if !result.Passed() {
		color = 31
		status = "FAIL"
	}

	// TODO: Check compatibility
	termColorFormat := "\033[%dm[%s]\n%s\033[0m"

	return fmt.Sprintf(termColorFormat, color, status, output)
}
