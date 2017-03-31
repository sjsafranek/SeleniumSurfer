package main

import (
	"fmt"
	"net/http"
)

const HTTP_DEFAULT_PORT = 8080

var HTTP_PORT int = HTTP_DEFAULT_PORT

type HttpServer struct {
	Port int
}

func (self HttpServer) Start() {
	// Attach Http Hanlders
	router := Router()
	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// Start server
	Ligneous.Info("[HttpServer] Magic happens on port ", self.Port)

	bind := fmt.Sprintf(":%v", self.Port)
	// bind := fmt.Sprintf("0.0.0.0:%v", port)

	err := http.ListenAndServe(bind, router)
	if err != nil {
		Ligneous.Error(err)
		shutDown()
	}

}
