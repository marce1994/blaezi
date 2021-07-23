# blaezi

blaezi is a super-lightweight smoke testing tool with zero dependencies, written purely in Go.

blaezi can perform smoke tests on endpoints specified by the user. blaezi also has the ability to check the content of the endpoint.

🔥 **You can see blaezi in action [here.](https://github.com/burntcarrot/blaezi-action-test/actions)**

## Installation

- **Using GitHub Actions:**
  - You can integrate blaezi with GitHub Actions. [Click here for an example.](https://github.com/burntcarrot/blaezi-action-test)
- **Using releases:**
  - Binaries are available under [Releases.](https://github.com/burntcarrot/blaezi/releases)
- **Using source code:**
  - You can clone the repo and run `go run main.go` for testing `blaezi`:
  - `git clone https://github.com/burntcarrot/blaezi && cd blaezi`
  - `go run main.go`


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
