package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// basic flags
	var testFile = flag.String("tests", "test.yml", "Complete path to the tests.")
	var secure = flag.Bool("secure", false, "Secure connection.")

	// parse flags and get args
	flag.Parse()
	args := flag.Args()

	// get the base URL at least
	if len(args) < 1 {
		fmt.Println("No base URL specified. Stopping smoke test.")
		os.Exit(0)
	}

	// throw err if the base URL is invalid
	baseURL, err := parseBaseURL(args[0])
	if err != nil {
		fmt.Println("Invalid base URL specified. Stopping smoke test.")
		os.Exit(0)
	}

	// extract tests from file
	extractedTests, err := extractTests(*testFile)
	if err != nil {
		fmt.Printf("Error raised while extracting tests: %s\n", err.Error())
		os.Exit(1)
	}

	if *secure == true {
		fmt.Println("secure mode.")
	} else {
		fmt.Println("insecure mode.")
	}

	// TODO:
	// create a new HTTP client
	client := Client(*secure)

	// create a new inspector object
	inspector := Inspector(client, baseURL)

	// process using the object
	results, errors := inspector.Test(extractedTests)

	// print result for each test
	for _, result := range results {
		fmt.Println(result)
	}

	// throw errors if raised (in tests)
	for _, err := range errors {
		fmt.Println(err)
	}

	// Show final stats (prettified)
	fmt.Println(results.countSuccess(), " tests passed out of", len(extractedTests), "tests.")
	// TODO:
	// Show time stats
	// Show failure message if tests fail
}
