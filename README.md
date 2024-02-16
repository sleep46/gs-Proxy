
# gs-Proxy

gs-Proxy is a simple Redis proxy that provides an HTTP API and caching layer for Redis GET commands.

## Features

- HTTP API: Make Redis GET command calls through a HTTP GET request.
- Cached GET: GET command requests are cached in an in-memory LRU cache.
- Configuration options: Customize cache capacity, global expiry time, Redis address, and server port.

## Usage

### Prerequisites

- Go (Golang) installed
- Docker and docker-compose installed (optional, for running in a container)

### Installation

To install gs-Proxy, clone the repository and build the binary:

```bash
git clone https://github.com/sleep46/gs-Proxy.git
cd gs-Proxy
make build
```

### Configuration

Configuration options can be set using environment variables:

- `CACHE_CAPACITY`: The maximum number of keys the cache retains.
- `GLOBAL_EXPIRY`: Expiry time for cache entries (in milliseconds).
- `PORT`: Port number for the HTTP server.
- `REDIS_ADDRESS`: Address of the Redis instance (in the format "host:port").

### Running

To run gs-Proxy, use the following command:

```bash
./gs-proxy
```

### Example

```go

package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    resp, err := http.Get("http://localhost:8080/GET/mykey")
    if err != nil {
        fmt.Println("GET request failed:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("response body failed:", err)
        return
    }

    fmt.Println("Response:", string(body))
}
```

## Contributing

Contributions are welcome just make a request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

