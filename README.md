# blaezi

blaezi is a super-lightweight smoke testing tool with zero dependencies, written purely in Go.

blaezi can perform smoke tests on endpoints specified by the user. blaezi also has the ability to check the content of the endpoint.

## Installation

blaezi is under testing for GitHub Actions. The README will be updated when tests are done and a release is made.

You can clone the repo and run `go run main.go` for testing `blaezi`.


## Usage

```
Usage of ./blaezi:
  -secure
        Secure connection.
  -tests string
        Complete path to the tests. (default "test.json")
  -timeout int
        Timeout for client. (default 5)
```

## Tests:

A test can be specified using JSON. `url` and `status_code` are necessary, `content` is optional.

Here is a sample test:

```json
[
    {
      "url": "/",
      "status_code": 200,
      "content": "Hello!"
    },
    {
      "url": "/posts",
      "status_code": 200,
      "content": "Hello post!"
    },
    {
      "url": "/nonexistent",
      "status_code": 404
    }
]
```


## Example

With the test specified above we can run `blaezi` like:

```
./blaezi http://localhost:8080 -tests test.json
```

If you want to specify a timeout for 10 seconds, you can do so by:

```
./blaezi http://localhost:8080 -tests test.json -timeout 10
```
