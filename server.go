package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"go.pixelfactory.io/pkg/observability/log"
	"go.pixelfactory.io/pkg/observability/log/fields"
)

// Stop handles OS Signals.
var (
	stop   = make(chan struct{})
	stopCh = setupSignalHandler(stop)
)

// Server is a http server.
type Server struct {
	Name                      string
	Router                    http.Handler
	Logger                    log.Logger
	Port                      string
	HTTPServerTimeout         time.Duration
	HTTPServerShutdownTimeout time.Duration
	TLSConfig                 *tls.Config
}

// Option is an option for New server.
type Option func(*Server)

// WithName set server name.
func WithName(n string) Option {
	return func(s *Server) {
		s.Name = n
	}
}

// WithRouter set server http Handler.
func WithRouter(r http.Handler) Option {
	return func(s *Server) {
		s.Router = r
	}
}

// WithLogger set server logger.
func WithLogger(l log.Logger) Option {
	return func(s *Server) {
		s.Logger = l
	}
}

// WithPort set server port.
func WithPort(p string) Option {
	return func(s *Server) {
		s.Port = p
	}
}

// WithHTTPServerTimeout set server http timeout.
func WithHTTPServerTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.HTTPServerTimeout = t
	}
}

// WithHTTPServerShutdownTimeout set server http shutdown timeout.
func WithHTTPServerShutdownTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.HTTPServerShutdownTimeout = t
	}
}

// WithTLSConfig set server tls.Config.
func WithTLSConfig(cfg *tls.Config) Option {
	return func(s *Server) {
		s.TLSConfig = cfg
	}
}

// New create new Server with default values.
func New(opts ...Option) (*Server, error) {
	// setup default server
	srv := &Server{
		Name:                      "default",
		Router:                    http.NewServeMux(),
		Port:                      "8080",
		HTTPServerTimeout:         60 * time.Second,
		HTTPServerShutdownTimeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(srv)
	}

	// setup default logger
	if srv.Logger == nil {
		srv.Logger = log.New().With(fields.String("name", srv.Name), fields.String("port", srv.Port))
		srv.Logger.Info("using default logger")
	}

	return srv, nil
}

// ListenAndServe start server.
func (s *Server) ListenAndServe() error {
	srv := &http.Server{
		Addr:         ":" + s.Port,
		Handler:      s.Router,
		WriteTimeout: s.HTTPServerTimeout,
		ReadTimeout:  s.HTTPServerTimeout,
		IdleTimeout:  2 * s.HTTPServerTimeout,
	}

	// Create listener
	var ln net.Listener
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Port))
	if err != nil {
		s.Logger.Error("Unable to create net.Listener", fields.Error(err))
		return err
	}

	if s.TLSConfig != nil {
		ln = tls.NewListener(ln, s.TLSConfig)
	}

	// run server in background
	go func() {
		s.Logger.Info("Starting server")
		if err := srv.Serve(ln); err != http.ErrServerClosed {
			s.Logger.Error("Server crashed", fields.Error(err))
		}
	}()

	// wait for SIGTERM or SIGINT
	<-stopCh
	ctx, cancel := context.WithTimeout(context.Background(), s.HTTPServerShutdownTimeout)
	defer cancel()

	s.Logger.Info("Shutting down server")

	// attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		s.Logger.Error("Server graceful shutdown failed", fields.Error(err))
		return err
	}

	s.Logger.Info("Server stopped")
	return nil
}

// Shutdown stops the server.
func (s *Server) Shutdown() {
	close(stop)
}
