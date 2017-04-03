package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	//http.Redirect(w, r, "Welcome to Selenium Surfer dudes!", 200)
	fmt.Fprintf(w, "Welcome to Selenium Surfer dudes!")
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

// NewTaskHandler
func NewTaskHandler(w http.ResponseWriter, r *http.Request) {
	// curl -H "Content-Type: application/json" -X POST -d '{"search":"paul ryan is gay","probability":3}' http://localhost:7777/api/v1/task

	// Get request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		message := fmt.Sprintf(" %v %v [500]", r.Method, r.URL.Path)
		Ligneous.Critical("[HttpServer]", r.RemoteAddr, message)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	r.Body.Close()

	// Unmarshal feature
	var task Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		message := fmt.Sprintf(" %v %v [400]", r.Method, r.URL.Path)
		Ligneous.Critical("[HttpServer]", r.RemoteAddr, message)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	DB.Add(task.Search, task.Probability)
	WCPool.Add(task.Search)

	js, err := MarshalJsonFromString(w, r, `{"status":"ok"}`)
	if err != nil {
		return
	}
	SendJsonResponse(w, r, js)

}
