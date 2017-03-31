package main

import (
	"net/http"
)

type apiRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type apiRoutes []apiRoute

var routes = apiRoutes{

	// Web Client apiRoutes
	apiRoute{"Index", "GET", "/", IndexHandler},

	// Health check
	apiRoute{"Ping", "GET", "/ping", PingHandler},
}
