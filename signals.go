package server

import (
	"os"
	"os/signal"
	"syscall"
)

//nolint:gochecknoglobals // Required for ensuring single signal handler
var (
	onlyOneSignalHandler = make(chan struct{})
	shutdownSignals      = []os.Signal{os.Interrupt, syscall.SIGTERM}
)

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func setupSignalHandler(stop chan struct{}) <-chan struct{} {
	close(onlyOneSignalHandler) // panics when called twice

	c := make(chan os.Signal, signalChannelBufferSize)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
