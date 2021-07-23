package main

type Test struct {
	URL                string
	StatusCodeExpected int
	Content            string
}

type Result struct {
	Test       Test
	StatusCode int
	Content    []byte
}

type Results []Result

// Returns true if the test passes successfully.
func (result Result) Passed() bool {
	// Status code should be checked first.
	if result.Test.StatusCodeExpected != result.StatusCode {
		return false
	}

	// Then, check the content.
	// TODO: Add string matching (i.e. if a part of string is equal, then tests should return true)
	if result.Test.Content != "" && result.Test.Content != string(result.Content) {
		return false
	}

	return true
}

// Return total count of tests that passed.
func (results Results) countSuccess() int {
	var count int

	for _, result := range results {
		if result.Passed() {
			count++
		}
	}

	return count
}
