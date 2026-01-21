package server_test

import (
	"crypto/tls"
	"net/http"
	"testing"
	"time"

	"go.pixelfactory.io/pkg/observability/log"

	"go.pixelfactory.io/pkg/server"
)

// FuzzWithName tests the WithName option with arbitrary strings.
func FuzzWithName(f *testing.F) {
	// Add seed corpus
	f.Add("test-server")
	f.Add("my-api")
	f.Add("")
	f.Add("server-123")
	f.Add("with spaces")
	f.Add("with\nnewline")
	f.Add("特殊字符")

	f.Fuzz(func(t *testing.T, name string) {
		srv, err := server.New(server.WithName(name))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if srv == nil {
			t.Fatal("server should not be nil")
		}
		if srv.Name != name {
			t.Errorf("expected name %q, got %q", name, srv.Name)
		}
	})
}

// FuzzWithPort tests the WithPort option with arbitrary port strings.
func FuzzWithPort(f *testing.F) {
	// Add seed corpus
	f.Add("8080")
	f.Add("3000")
	f.Add("80")
	f.Add("443")
	f.Add("65535")
	f.Add("0")

	f.Fuzz(func(t *testing.T, port string) {
		srv, err := server.New(server.WithPort(port))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if srv == nil {
			t.Fatal("server should not be nil")
		}
		if srv.Port != port {
			t.Errorf("expected port %q, got %q", port, srv.Port)
		}
	})
}

// FuzzWithHTTPServerTimeout tests the WithHTTPServerTimeout option.
func FuzzWithHTTPServerTimeout(f *testing.F) {
	// Add seed corpus
	f.Add(int64(60))
	f.Add(int64(30))
	f.Add(int64(0))
	f.Add(int64(120))
	f.Add(int64(1))
	f.Add(int64(-1))

	f.Fuzz(func(t *testing.T, seconds int64) {
		timeout := time.Duration(seconds) * time.Second
		srv, err := server.New(server.WithHTTPServerTimeout(timeout))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if srv == nil {
			t.Fatal("server should not be nil")
		}
		if srv.HTTPServerTimeout != timeout {
			t.Errorf("expected timeout %v, got %v", timeout, srv.HTTPServerTimeout)
		}
	})
}

// FuzzWithHTTPServerShutdownTimeout tests the WithHTTPServerShutdownTimeout option.
func FuzzWithHTTPServerShutdownTimeout(f *testing.F) {
	// Add seed corpus
	f.Add(int64(10))
	f.Add(int64(5))
	f.Add(int64(0))
	f.Add(int64(30))
	f.Add(int64(1))
	f.Add(int64(-1))

	f.Fuzz(func(t *testing.T, seconds int64) {
		timeout := time.Duration(seconds) * time.Second
		srv, err := server.New(server.WithHTTPServerShutdownTimeout(timeout))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if srv == nil {
			t.Fatal("server should not be nil")
		}
		if srv.HTTPServerShutdownTimeout != timeout {
			t.Errorf("expected shutdown timeout %v, got %v", timeout, srv.HTTPServerShutdownTimeout)
		}
	})
}

// FuzzNew tests the New function with multiple options.
func FuzzNew(f *testing.F) {
	// Add seed corpus
	f.Add("test-server", "8080", int64(60), int64(10))
	f.Add("api", "3000", int64(30), int64(5))
	f.Add("", "80", int64(0), int64(0))

	f.Fuzz(func(t *testing.T, name, port string, timeout, shutdownTimeout int64) {
		timeoutDuration := time.Duration(timeout) * time.Second
		shutdownTimeoutDuration := time.Duration(shutdownTimeout) * time.Second

		srv, err := server.New(
			server.WithName(name),
			server.WithPort(port),
			server.WithHTTPServerTimeout(timeoutDuration),
			server.WithHTTPServerShutdownTimeout(shutdownTimeoutDuration),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if srv == nil {
			t.Fatal("server should not be nil")
		}

		// Verify all options were applied
		if srv.Name != name {
			t.Errorf("expected name %q, got %q", name, srv.Name)
		}
		if srv.Port != port {
			t.Errorf("expected port %q, got %q", port, srv.Port)
		}
		if srv.HTTPServerTimeout != timeoutDuration {
			t.Errorf("expected timeout %v, got %v", timeoutDuration, srv.HTTPServerTimeout)
		}
		if srv.HTTPServerShutdownTimeout != shutdownTimeoutDuration {
			t.Errorf("expected shutdown timeout %v, got %v", shutdownTimeoutDuration, srv.HTTPServerShutdownTimeout)
		}
	})
}

// FuzzWithRouter tests the WithRouter option.
func FuzzWithRouter(f *testing.F) {
	// Add seed corpus
	f.Add("")

	f.Fuzz(func(t *testing.T, _ string) {
		// Create a new router for each iteration
		router := http.NewServeMux()
		srv, err := server.New(server.WithRouter(router))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if srv == nil {
			t.Fatal("server should not be nil")
		}
		if srv.Router == nil {
			t.Fatal("router should not be nil")
		}
	})
}

// FuzzWithLogger tests the WithLogger option.
func FuzzWithLogger(f *testing.F) {
	// Add seed corpus
	f.Add("")

	f.Fuzz(func(t *testing.T, _ string) {
		// Create a new logger for each iteration
		logger := log.New()
		srv, err := server.New(server.WithLogger(logger))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if srv == nil {
			t.Fatal("server should not be nil")
		}
		if srv.Logger == nil {
			t.Fatal("logger should not be nil")
		}
	})
}

// FuzzWithTLSConfig tests the WithTLSConfig option.
func FuzzWithTLSConfig(f *testing.F) {
	// Add seed corpus with different MinVersion values
	f.Add(uint16(tls.VersionTLS10))
	f.Add(uint16(tls.VersionTLS11))
	f.Add(uint16(tls.VersionTLS12))
	f.Add(uint16(tls.VersionTLS13))

	f.Fuzz(func(t *testing.T, minVersion uint16) {
		// Create a basic TLS config (without loading certificates to avoid file I/O)
		tlsConfig := &tls.Config{
			MinVersion: minVersion,
		}

		srv, err := server.New(server.WithTLSConfig(tlsConfig))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if srv == nil {
			t.Fatal("server should not be nil")
		}
		if srv.TLSConfig == nil {
			t.Fatal("TLS config should not be nil")
		}
		if srv.TLSConfig.MinVersion != minVersion {
			t.Errorf("expected MinVersion %v, got %v", minVersion, srv.TLSConfig.MinVersion)
		}
	})
}
