package main

import (
	"fmt"
	"net/http"
	"time"

	"go.pixelfactory.io/pkg/server"
	"go.uber.org/zap"
)

func main() {

	logger := zap.NewExample()
	defer logger.Sync()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	// Start http server
	srv, err := server.NewServer(
		server.WithName("Web"),
		server.WithRouter(mux),
		server.WithLogger(logger),
		server.WithConfig(&server.Config{
			Port:                      "3000",
			HTTPServerTimeout:         10 * time.Second,
			HTTPServerShutdownTimeout: 5 * time.Second,
		}),
	)
	if err != nil {
		logger.Fatal("Unable to create server", zap.Error(err))
	}

	// This will block until shutdown
	srv.ListenAndServe()
}
