package main

import "net/url"

type Inspector struct {
	client  Client
	baseURL *url.URL
}

func NewInspector(client Client, baseURL *url.URL) Inspector {
	return Inspector{client: client, baseURL: baseURL}
}
