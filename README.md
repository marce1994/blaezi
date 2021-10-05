# 🔥 blaezi
# to test

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
- **Docker:**
    You can build the Docker image using:
  - `git clone https://github.com/burntcarrot/blaezi`
  - `docker build . -t blaezi -f Dockerfile --no-cache`
  - and run it
  ```
  docker run \
    --rm \
    -t \
    -v "$PWD/YOURTEST.json":"/test.json" \
   -it IMAGE_ID https://someurl.io -tests "test.json"
  ```

## Usage

```
Usage of ./blaezi:
  -auth string
        Authorization string.
  -secure
        Secure connection.
  -tests string
        Complete path to the tests file. (default "test.json")
  -timeout int
        Timeout for client. (default 5)
```

## Tests:

A test can be specified using JSON. `url`, `status_code` and `method` are necessary, `content`, `request_body` are optional.

Here is a sample test:

```json
[
  {
    "url": "/",
    "status_code": 200,
    "content": "Hello",
    "method": "GET"
  },
  {
    "url": "/posts",
    "status_code": 200,
    "content": "Hello",
    "method": "GET"
  },
  {
    "url": "/posts",
    "status_code": 200,
    "content": "Hello",
    "method": "POST",
    "request_body": "{\"id\": 1, \"title\": \"H\", \"text\": \"Hello\"}"
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

#### Acknowledgements:

Some parts of `blaezi` were inspired by [vape](https://github.com/symm/vape).

`blaezi` is rewritten from scratch and aims to add additional features on top of [vape](https://github.com/symm/vape). (Spoiler: `blaezi` currently provides some features like GitHub integration, etc.)
