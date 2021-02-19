package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"go.pixelfactory.io/pkg/server"
)

func main() {

	logger := zap.NewExample()
	defer logger.Sync()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Start http server
	srv, err := server.NewServer(
		server.WithName("Web"),
		server.WithRouter(mux),
		server.WithLogger(logger),
		server.WithPort("3000"),
		server.WithTLSConfig(&tls.Config{Certificates: []tls.Certificate{cer}}),
	)
	if err != nil {
		logger.Fatal("Unable to create server", zap.Error(err))
	}

	// This will block until shutdown
	srv.ListenAndServe()
}
