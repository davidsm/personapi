package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masenius/personapi/person"
)

type Options struct {
	Seed *int64
}

func Create(opts *Options) http.Handler {
	if opts.Seed != nil {
		person.Seed(*opts.Seed)
	}

	router := httprouter.New()
	router.GET("/", handleRequest)
	return router
}
