package main

import (
	"net/http"
	"runtime"
	"time"
)

var serverStartTime time.Time

func init() {
	serverStartTime = time.Now()
}

// IndexHandler returns html page containing api docs
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://sjsafranek.github.io/gospatial/", 200)
	return
}

// PingHandler provides an api route for server health check
func PingHandler(w http.ResponseWriter, r *http.Request) {
	Ligneous.Debug("[HttpServer] ", r)

	var data map[string]interface{}
	data = make(map[string]interface{})
	data["status"] = "success"
	result := make(map[string]interface{})
	result["result"] = "pong"
	result["registered"] = serverStartTime.UTC()
	result["uptime"] = time.Since(serverStartTime).Seconds()
	result["num_cores"] = runtime.NumCPU()
	data["data"] = result

	js, err := MarshalJsonFromStruct(w, r, data)
	if err != nil {
		return
	}

	SendJsonResponse(w, r, js)
}
