package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"go.pixelfactory.io/pkg/server"
)

func main() {
	// Initialize logger
	logger := zap.NewExample()
	defer func() {
		_ = logger.Sync() // Flush buffered logs
	}()

	// Setup HTTP routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/api/v1/users", handleUsers)

	// Load TLS certificates
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		logger.Fatal("Failed to load TLS certificates", zap.Error(err))
	}

	// Create server with all options
	srv, err := server.New(
		server.WithName("ExampleAPI"),
		server.WithRouter(mux),
		server.WithLogger(logger),
		server.WithPort("3000"),
		server.WithTLSConfig(&tls.Config{
			Certificates: []tls.Certificate{cer},
			MinVersion:   tls.VersionTLS13,
		}),
	)
	if err != nil {
		logger.Fatal("Failed to create server", zap.Error(err))
	}

	logger.Info("Server configured, starting...")

	// ListenAndServe blocks until SIGTERM/SIGINT is received
	// It handles graceful shutdown automatically
	if serveErr := srv.ListenAndServe(); serveErr != nil {
		logger.Fatal("Server failed", zap.Error(serveErr))
	}

	logger.Info("Server shutdown complete")
}

// handleRoot handles requests to the root path.
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// handleHealth provides a health check endpoint.
func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status":"healthy"}`)
}

// handleUsers demonstrates an API endpoint.
func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"users":["alice","bob"]}`)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
