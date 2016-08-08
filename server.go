package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/masenius/personapi/app"
	"net/http"
)

const defaultPort = 8080

func main() {
	port := flag.Int("port", defaultPort, fmt.Sprintf("Port to use. Defaults to %d", defaultPort))
	bind := flag.String("bind", "", "Bind to address. Default is empty, meaning 0.0.0.0")
	flag.Parse()
	address := fmt.Sprintf("%s:%d", *bind, *port)

	router := httprouter.New()
	router.GET("/", app.HandleRequest)

	fmt.Println("Starting server on", address)
	http.ListenAndServe(address, router)
}
