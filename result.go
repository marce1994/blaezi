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

import "strings"

type Test struct {
	URL                string `json:"url"`
	StatusCodeExpected int    `json:"status_code"`
	Content            string `json:"content"`
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
	if result.Test.Content != "" && !strings.Contains(string(result.Content), result.Test.Content) {
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
