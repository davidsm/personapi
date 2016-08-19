package app

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masenius/personapi/person"
)

type Options struct {
	Seed   *int64
	Logger *log.Logger
}

func Create(opts *Options) http.Handler {
	if opts.Seed != nil {
		person.Seed(*opts.Seed)
	}

	router := httprouter.New()
	router.GET("/", handleRequest)
	router.HEAD("/", handleRequest)

	if opts.Logger == nil {
		return router
	}

	return &logHandler{router, opts.Logger}
}
