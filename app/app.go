package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Create() http.Handler {
	router := httprouter.New()
	router.GET("/", HandleRequest)
	return router
}
