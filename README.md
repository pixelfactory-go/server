# Server Package

[![Go Report Card](https://goreportcard.com/badge/go.pixelfactory.io/pkg/server)](https://goreportcard.com/report/go.pixelfactory.io/pkg/server)
[![Go Reference](https://pkg.go.dev/badge/go.pixelfactory.io/pkg/server.svg)](https://pkg.go.dev/go.pixelfactory.io/pkg/server)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lightweight, production-ready HTTP server package for Go applications with built-in graceful shutdown, TLS support, and structured logging.

## Features

- **Graceful Shutdown**: Automatic handling of SIGTERM and SIGINT signals with configurable shutdown timeout
- **TLS Support**: Optional TLS configuration for secure connections
- **Structured Logging**: Integrated with pixelfactory observability logging framework
- **Configurable Timeouts**: Customizable read, write, and idle timeouts
- **Functional Options**: Clean API using the functional options pattern
- **Production Ready**: Battle-tested with sensible defaults

## Installation

```bash
go get go.pixelfactory.io/pkg/server
```

## Quick Start

### Basic HTTP Server

```go
package main

import (
    "net/http"
    "go.pixelfactory.io/pkg/server"
)

func main() {
    // Create a simple router
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    // Create and start server with default settings
    srv, err := server.New(
        server.WithRouter(mux),
    )
    if err != nil {
        panic(err)
    }

    // Start server (blocks until shutdown signal received)
    if err := srv.ListenAndServe(); err != nil {
        panic(err)
    }
}
```

### Customized Server

```go
package main

import (
    "net/http"
    "time"
    "go.pixelfactory.io/pkg/server"
    "go.pixelfactory.io/pkg/observability/log"
)

func main() {
    // Create custom logger
    logger := log.New()

    // Create router
    mux := http.NewServeMux()
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })

    // Create server with custom configuration
    srv, err := server.New(
        server.WithName("my-api-server"),
        server.WithRouter(mux),
        server.WithLogger(logger),
        server.WithPort("3000"),
        server.WithHTTPServerTimeout(30*time.Second),
        server.WithHTTPServerShutdownTimeout(15*time.Second),
    )
    if err != nil {
        panic(err)
    }

    // Start server
    if err := srv.ListenAndServe(); err != nil {
        panic(err)
    }
}
```

### TLS/HTTPS Server

```go
package main

import (
    "crypto/tls"
    "net/http"
    "go.pixelfactory.io/pkg/server"
)

func main() {
    // Load TLS certificates
    cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
    if err != nil {
        panic(err)
    }

    tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{cert},
        MinVersion:   tls.VersionTLS12,
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Secure Hello!"))
    })

    // Create server with TLS
    srv, err := server.New(
        server.WithRouter(mux),
        server.WithPort("8443"),
        server.WithTLSConfig(tlsConfig),
    )
    if err != nil {
        panic(err)
    }

    if err := srv.ListenAndServe(); err != nil {
        panic(err)
    }
}
```

## Configuration Options

The server can be configured using functional options:

| Option | Description | Default |
|--------|-------------|---------|
| `WithName(string)` | Set server name for logging | `"default"` |
| `WithRouter(http.Handler)` | Set HTTP router/handler | `http.NewServeMux()` |
| `WithLogger(log.Logger)` | Set custom logger | Default logger |
| `WithPort(string)` | Set server port | `"8080"` |
| `WithHTTPServerTimeout(time.Duration)` | Set read/write timeout | `60s` |
| `WithHTTPServerShutdownTimeout(time.Duration)` | Set graceful shutdown timeout | `10s` |
| `WithTLSConfig(*tls.Config)` | Enable TLS with configuration | `nil` (disabled) |

## Graceful Shutdown

The server automatically handles OS signals (SIGTERM, SIGINT) for graceful shutdown:

1. Signal received → Server stops accepting new connections
2. Existing requests complete within shutdown timeout
3. Server logs shutdown events
4. Clean exit

You can also programmatically trigger shutdown:

```go
srv, _ := server.New()

// Start server in goroutine
go func() {
    srv.ListenAndServe()
}()

// Trigger shutdown from your code
srv.Shutdown()
```

## Default Behavior

When created with `New()` without options:

- **Name**: `"default"`
- **Port**: `8080`
- **Router**: Empty `http.NewServeMux()`
- **Timeouts**: 60s read/write, 120s idle
- **Shutdown Timeout**: 10s
- **Logger**: Default structured logger with server name and port fields
- **TLS**: Disabled

## Development

### Prerequisites

- Go 1.24 or higher
- golangci-lint (for linting)

### Running Tests

```bash
make test
```

### Linting

```bash
make lint
```

### Formatting

```bash
make fmt
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Related Projects

- [pixelfactory-go/version](https://github.com/pixelfactory-go/version) - Build-time version information
- [pixelfactory-go/observability](https://github.com/pixelfactory-go/observability) - Structured logging and observability

## Support

For issues, questions, or contributions, please use the [GitHub issue tracker](https://github.com/pixelfactory-go/server/issues).
