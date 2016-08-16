package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/masenius/personapi/app"
)

const defaultPort = 8080

func main() {
	port := flag.Int("port", defaultPort, fmt.Sprintf("Port to use. Defaults to %d", defaultPort))
	bind := flag.String("bind", "", "Bind to address. Default is empty, meaning 0.0.0.0")
	seed := flag.Int64("seed", 0, "Specify seed for the random generator. 0 means seed with current time. Not including this argument has the same effect as 0")
	flag.Parse()
	address := fmt.Sprintf("%s:%d", *bind, *port)

	var seedOpt *int64
	if *seed != 0 {
		seedOpt = seed
	}
	appOptions := app.Options{
		Seed: seedOpt,
	}
	app := app.Create(&appOptions)

	fmt.Println("Starting server on", address)
	http.ListenAndServe(address, app)
}
