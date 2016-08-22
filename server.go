package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/masenius/personapi/app"
	"github.com/masenius/personapi/reqlog"
)

const defaultPort = 8080

func main() {
	port := flag.Int("port", defaultPort, fmt.Sprintf("Port to use"))
	bind := flag.String("bind", "", "Bind to address. Default is empty, meaning 0.0.0.0")
	seed := flag.Int64("seed", 0, "Specify seed for the random generator. 0 means seed with current time. Not including this argument has the same effect as 0")
	logFile := flag.String("logfile", "", "Log to file path. If not specified, log to stdout")
	flag.Parse()
	address := fmt.Sprintf("%s:%d", *bind, *port)

	var seedOpt *int64
	if *seed != 0 {
		seedOpt = seed
	}

	var logger *log.Logger
	if *logFile == "" {
		logger = reqlog.Stdout()
	} else {
		var err error
		logger, err = reqlog.File(*logFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	appOptions := app.Options{
		Seed:   seedOpt,
		Logger: logger,
	}
	app := app.Create(&appOptions)

	logger.Println("Starting server on", address)
	logger.Fatal(http.ListenAndServe(address, app))
}
