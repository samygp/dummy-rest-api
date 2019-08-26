package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/samygp/dummy-rest-api/config"
)

func newServer() *http.Server {
	handler := NewHandler()
	router := NewRouter(handler)
	// Create a new server and set timeout values.
	return &http.Server{
		Addr:           ":" + config.Config.Server.Port,
		Handler:        router,
		ReadTimeout:    time.Duration(config.Config.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Config.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

//Start creates a server and starts listening
func Start() {
	srv := newServer()

	// We want to report the listener is closed.
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		log.Infof("%s server is now running on port %s", config.Config.App.Name, config.Config.Server.Port)
		srv.ListenAndServe()
		wg.Done()
	}()

	// Listen for an interrupt signal from the OS.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt)

	// Wait for a signal to shutdown.
	<-osSignals

	closeServer(srv)

	// Wait for the listener to report it is closed.
	wg.Wait()
}

//closeServer shuts down the server and tries to finish all pending operations
func closeServer(srv *http.Server) {
	// Create a context to attempt a graceful 5 second shutdown.
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Attempt the graceful shutdown by closing the listener and
	// completing all inflight requests.
	if err := srv.Shutdown(ctx); err != nil {

		log.Errorf("Error shutting down server: %s", err)

		// Looks like we timedout on the graceful shutdown. Kill it hard.
		if err := srv.Close(); err != nil {
			log.Errorf("Error killing server: %s", err)
		}
	}
}
