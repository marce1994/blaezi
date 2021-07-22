package main

import (
	"crypto/tls"
	"net/http"
	"time"
)

// the client should be able to perform a request
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
