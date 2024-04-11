package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"template-app/cmd/log"
	"template-app/src/handler"
	"template-app/src/middleware"
	"template-app/src/router"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Create(log *log.Logger) *mux.Router {
	log.System("Server routes building...")
	muxRouter := mux.NewRouter()
	muxRouter.Use(middleware.LogRequest(log))
	// /row route builder
	(&router.Product{
		Log:     log,
		Router:  muxRouter.PathPrefix("/product/").Subrouter(),
		Handler: &handler.Product{Log: log},
	}).Build()
	// /swagger route builder
	(&router.Swagger{
		Log:    log,
		Router: muxRouter,
	}).Build()

	return muxRouter
}

func startServer(log *log.Logger) *http.Server {

	muxRouter := Create(log)
	port := "3000"
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      muxRouter,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	go func() {
		log.System("Starting server... in port: %s", port)
		err := httpServer.ListenAndServe()
		if err != nil {
			log.SystemFatal("Error starting server: %v", err)
		}
	}()

	return httpServer
}

func waitForShutdown(s *http.Server, l *log.Logger) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	sig := <-sigChan
	l.System("Received terminate, graceful shutdown", sig)
	d := 30 * time.Second
	tc, cancel := context.WithTimeout(context.Background(), d)
	s.Shutdown(tc)
	defer cancel()
}

func startEnv(log *log.Logger) {
	if os.Getenv("APP_ENV") == "" {
		log.System("Loading .env file...")
		err := godotenv.Load(".env")
		if err != nil {
			err := godotenv.Load("../.env")
			if err != nil {
				log.System("Error loading .env file: %v", err)
			}
		}
	} else {
		os.Setenv("APP_ENV", "Production")
	}
	log.System("Environment: %s", os.Getenv("APP_ENV"))
}
