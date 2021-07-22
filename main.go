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
	// create a new object

	// process using the object
	// get result

	// print result for each test
	// throw errors if raised (in tests)

	// Show final stats (prettified)
}
