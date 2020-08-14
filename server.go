package server

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Handle OS Signals
var stopCh = setupSignalHandler()

// Config holds server config
type Config struct {
	Port                      string        `mapstructure:"port"`
	HTTPServerTimeout         time.Duration `mapstructure:"http-server-timeout"`
	HTTPServerShutdownTimeout time.Duration `mapstructure:"http-server-shutdown-timeout"`
}

// Logger is a simplified abstraction of the zap.Logger
type Logger interface {
	Info(msg string, fields ...zapcore.Field)
	Warn(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
}

// Server holds server
type Server struct {
	name   string
	router http.Handler
	config *Config
	logger Logger
}

// Option type
type Option func(*Server)

// WithName set server name
func WithName(n string) Option {
	return func(s *Server) {
		s.name = n
	}
}

// WithRouter set server http Handler
func WithRouter(r http.Handler) Option {
	return func(s *Server) {
		s.router = r
	}
}

// WithConfig set server config
func WithConfig(c *Config) Option {
	return func(s *Server) {
		s.config = c
	}
}

// WithLogger set server logger
func WithLogger(l Logger) Option {
	return func(s *Server) {
		s.logger = l
	}
}

// NewServer create new Server with default values
func NewServer(opts ...Option) (*Server, error) {
	srv := &Server{
		name:   "default",
		router: http.NewServeMux(),
		config: &Config{
			Port:                      "8080",
			HTTPServerTimeout:         60 * time.Second,
			HTTPServerShutdownTimeout: 5 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(srv)
	}

	return srv, nil
}

// ListenAndServe start server
func (s *Server) ListenAndServe() {

	srv := &http.Server{
		Addr:         ":" + s.config.Port,
		WriteTimeout: s.config.HTTPServerTimeout,
		ReadTimeout:  s.config.HTTPServerTimeout,
		IdleTimeout:  2 * s.config.HTTPServerTimeout,
		Handler:      s.router,
	}

	// run server in background
	go func() {
		s.logger.Info("Starting HTTP server", zap.String("name", s.name), zap.String("port", s.config.Port))
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Fatal("HTTP server crashed", zap.String("name", s.name), zap.Error(err))
		}
	}()

	// wait for SIGTERM or SIGINT
	<-stopCh
	ctx, cancel := context.WithTimeout(context.Background(), s.config.HTTPServerShutdownTimeout)
	defer cancel()

	s.logger.Info("Shutting down HTTP server", zap.String("name", s.name), zap.Duration("timeout", s.config.HTTPServerShutdownTimeout))

	// attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		s.logger.Warn("HTTP server graceful shutdown failed", zap.String("name", s.name), zap.Error(err))
	} else {
		s.logger.Info("HTTP server stopped", zap.String("name", s.name))
	}
}
