package server_test

import (
	"crypto/tls"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.pixelfactory.io/pkg/observability/log"

	"go.pixelfactory.io/pkg/server"
)

func Test_NewServer(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	httpSrv, err := server.New()
	is.NoError(err)
	is.NotEmpty(httpSrv)
	is.Equal(httpSrv.Port, "8080")
	is.Equal(httpSrv.HTTPServerTimeout, 60*time.Second)
	is.Equal(httpSrv.HTTPServerShutdownTimeout, 10*time.Second)
}

func Test_NewServer_WithName(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	httpSrv, err := server.New(
		server.WithName("test"),
	)
	is.NoError(err)
	is.NotEmpty(httpSrv)
	is.Equal(httpSrv.Name, "test")
}

func Test_NewServer_WithPort(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	httpSrv, err := server.New(
		server.WithPort("1234"),
	)
	is.NoError(err)
	is.NotEmpty(httpSrv)
	is.Equal(httpSrv.Port, "1234")
}

func Test_NewServer_WithLogger(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	logger := log.New()
	httpSrv, err := server.New(
		server.WithLogger(logger),
	)
	is.NoError(err)
	is.NotEmpty(httpSrv)
	is.Equal(httpSrv.Logger, logger)
}

func Test_NewServer_WithRouter(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	router := http.NewServeMux()
	httpSrv, err := server.New(
		server.WithRouter(router),
	)
	is.NoError(err)
	is.NotEmpty(httpSrv)
	is.Equal(httpSrv.Router, router)
}

func Test_NewServer_WithHTTPServerTimeout(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	httpSrv, err := server.New(
		server.WithHTTPServerTimeout(10 * time.Second),
	)
	is.NoError(err)
	is.NotEmpty(httpSrv)
	is.Equal(httpSrv.HTTPServerTimeout, 10*time.Second)
}

func Test_NewServer_WithHTTPServerShutdownTimeout(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	httpSrv, err := server.New(
		server.WithHTTPServerShutdownTimeout(10 * time.Second),
	)
	is.NoError(err)
	is.NotEmpty(httpSrv)
	is.Equal(httpSrv.HTTPServerShutdownTimeout, 10*time.Second)
}

func Test_NewServer_WithTLSConfig(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	cer, err := tls.LoadX509KeyPair("example/server.crt", "example/server.key")
	if err != nil {
		t.Fatal(err.Error())
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cer},
		MinVersion:   tls.VersionTLS12, // Set the minimum TLS version to TLS 1.2
	}

	httpSrv, err := server.New(
		server.WithTLSConfig(tlsConfig),
	)
	is.NoError(err)
	is.NotEmpty(httpSrv)
	is.Equal(httpSrv.TLSConfig, tlsConfig)
}

func Test_NewServer_ListenAndServe(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	cer, err := tls.LoadX509KeyPair("example/server.crt", "example/server.key")
	if err != nil {
		t.Fatal(err.Error())
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cer},
		MinVersion:   tls.VersionTLS12, // Set the minimum TLS version to TLS 1.2
	}

	httpSrv, err := server.New(
		server.WithTLSConfig(tlsConfig),
	)
	is.NoError(err)
	is.NotEmpty(httpSrv)

	serviceRunning := make(chan struct{})
	serviceDone := make(chan struct{})
	go func() {
		close(serviceRunning)
		err := httpSrv.ListenAndServe()
		is.NoError(err)
		defer close(serviceDone)
	}()

	// wait until the goroutine started to run (1)
	<-serviceRunning

	httpSrv.Shutdown()

	// wait until the service is shutdown (3)
	<-serviceDone
}

func Test_NewServer_ListenAndServe_Error(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	httpSrv, err := server.New(
		server.WithPort("invalid"),
	)
	is.NoError(err)
	is.NotEmpty(httpSrv)

	err = httpSrv.ListenAndServe()
	is.Error(err)
}
