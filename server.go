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
	flag.Parse()
	address := fmt.Sprintf("%s:%d", *bind, *port)

	app := app.Create()

	fmt.Println("Starting server on", address)
	http.ListenAndServe(address, app)
}
