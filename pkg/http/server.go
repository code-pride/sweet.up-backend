package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type HttpConfig struct {
	Host         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

func InitHttp(router *mux.Router, log *zap.SugaredLogger) {

	s := http.Server{
		Addr:         "localhost:8080",  // configure the bind address
		Handler:      router,            // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		log.Info("Starting server on port 8080")

		err := s.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Error("Error starting server: %s", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Info("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
