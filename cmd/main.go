package main

import (
	"context"
	"dimension/internal/api"
	"dimension/internal/logger"
	"dimension/internal/middleware"
	"dimension/internal/storage"
	"github.com/namsral/flag"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	foo                      = false
	restSrv     *http.Server = nil
	waitGroup   sync.WaitGroup
	serviceName string
)

const defaultPORT = "8080"

func main() {

	logger.Init()
	logger.Log.Error()

	serviceName = "Dimension"

	logger.Log.Info("Service/Instance name: " + serviceName)

	srv := &http.Server{Addr: defaultPORT}

	///////////////////////////////
	//ENVIRONMENT PARSING
	//////////////////////////////

	flag.BoolVar(&foo, "foo", false, "foo") //this was a flag useful for dev work

	flag.Parse()

	logger.Log.SetReportCaller(true)

	///////////////////////////////
	//CONFIGURATION
	//////////////////////////////

	var gameProvider storage.GameProvider
	gameProvider = storage.NewMemGame()

	middleware.GameProvider = gameProvider
	///////////////////////////////
	//SIGNAL HANDLING
	//////////////////////////////

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	idleConnsClosed := make(chan struct{})
	go func() {
		<-c
		foo = false

		//cancel()

		// Gracefully shutdown the HTTP server.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logger.Log.Infof("HTTP server Shutdown: %v", err)
		}

		ctx := context.Background()
		if restSrv != nil {
			if err := restSrv.Shutdown(ctx); err != nil {
				logger.Log.Panic("ERROR: Unable to stop REST collection server!")
			}
		}

		close(idleConnsClosed)
	}()

	///////////////////////
	// START
	///////////////////////

	//Rest Server
	waitGroup.Add(1)
	doRest(defaultPORT)

	//////////////////////
	// WAIT FOR GOD
	/////////////////////

	waitGroup.Wait()
	logger.Log.Info("HTTP server shutdown, waiting for idle connection to close...")
	<-idleConnsClosed
	logger.Log.Info("Done. Exiting.")
}

func doRest(serverPort string) {
	logger.Log.Info("**RUNNING -- Listening on " + defaultPORT)

	router := api.NewRouter()

	srv := &http.Server{
		Addr:    ":" + serverPort,
		Handler: router,
	}

	go func() {
		defer waitGroup.Done()
		if err := srv.ListenAndServe(); err != nil {
			// Cannot panic because this is probably just a graceful shutdown.
			logger.Log.Error(err)
			logger.Log.Info("REST collection server shutdown.")
		}
	}()

	logger.Log.Info("REST collection server started on port " + serverPort)
	restSrv = srv
}

//todo
//need to create an expose an API;
//probably need to create a VERY VERY VERY dumb AI; like only obeys the quantity; places it whereever (that way I can generate some more test data)
//need to think about creating a CLI to interact with it? can it remember game session identifier? etc?
//I think player and game needs tokens to protect them from unlawful submissions
