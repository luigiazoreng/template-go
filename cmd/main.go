package main

import (
	log "template-app/cmd/log"
)

// @title Template-Go Microservice
// @version 1.0
// @description API template
// @host localhost:3000
// @template-app /src
//@BasePath /

// It creates a new server and logs any errors.
func main() {
	log := log.StartLogger()
	startEnv(log)
	httpServer := startServer(log)
	waitForShutdown(httpServer, log)
}
