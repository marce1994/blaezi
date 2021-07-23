package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"sync"
)

type Inspector struct {
	client  Client
	baseURL *url.URL
}

func NewInspector(client Client, baseURL *url.URL) Inspector {
	return Inspector{client: client, baseURL: baseURL}
}

type Tests []Test

// The worker fetches tests and performs them.
func (inspector Inspector) worker(wg *sync.WaitGroup, testsChan <-chan Test, resultsChan chan<- Result, errorChan chan<- error) {
	for test := range testsChan {
		result, err := inspector.runTest(test)

		if err != nil {
			errorChan <- err
		} else {
			resultsChan <- result
		}

		wg.Done()
	}
}

// Runs tests and returns results.
func (inspector Inspector) runTest(test Test) (Result, error) {
	url := *inspector.baseURL

	// Parses URL, checks for invalid URLs
	parsedURL, err := url.Parse(path.Join(url.Path + test.URL))
	if err != nil {
		return Result{}, err
	}

	// Send a new GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return Result{}, err
	}

	// Get response
	resp, err := inspector.client.Do(req)
	if err != nil {
		return Result{}, err
	}

	result := Result{StatusCode: resp.StatusCode, Test: test}

	// If content is specified in test, read response body.
	if test.Content != "" {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return result, err
		}

		result.Content = content
	}

	return result, nil
}

// This is exposed to the main program.
func (inspector Inspector) Test(tests Tests) (results Results, errors []error) {
	// total count of tests
	numTests := len(tests)

	// Channels for handling tests, results and errors.
	testsChan := make(chan Test, numTests)
	resultsChan := make(chan Result, numTests)
	errorChan := make(chan error, numTests)

	var wg sync.WaitGroup

	// TODO: Make number of goroutines configurable
	for i := 1; i <= 4; i++ {
		// send workers to process tests
		go inspector.worker(&wg, testsChan, resultsChan, errorChan)
	}

	// Send tests to the tests channel.
	for _, test := range tests {
		testsChan <- test
		wg.Add(1)
	}

	// Close test channel
	close(testsChan)

	wg.Wait()

	// fetch and append results
	for i := 0; i < numTests; i++ {
		select {
		case err := <-errorChan:
			errors = append(errors, err)
		case result := <-resultsChan:
			results = append(results, result)
		}
	}

	return results, errors
}
