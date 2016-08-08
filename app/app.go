package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Create() http.Handler {
	router := httprouter.New()
	router.GET("/", HandleRequest)
	return router
}
