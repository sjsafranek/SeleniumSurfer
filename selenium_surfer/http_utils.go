package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MarshalJsonFromString(w http.ResponseWriter, r *http.Request, data string) ([]byte, error) {
	js, err := json.Marshal(data)
	if err != nil {
		message := fmt.Sprintf(" %v %v [500]", r.Method, r.URL.Path)
		Ligneous.Critical("[HttpServer]", r.RemoteAddr, message)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return js, err
	}
	return js, nil
}

func MarshalJsonFromStruct(w http.ResponseWriter, r *http.Request, data interface{}) ([]byte, error) {
	js, err := json.Marshal(data)
	if err != nil {
		message := fmt.Sprintf(" %v %v [500]", r.Method, r.URL.Path)
		Ligneous.Critical("[HttpServer]", r.RemoteAddr, message)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return js, err
	}
	return js, nil
}

// Sends http response
func SendJsonResponse(w http.ResponseWriter, r *http.Request, js []byte) {
	// Log result
	message := fmt.Sprintf(" %v %v [200]", r.Method, r.URL.Path)
	Ligneous.Info("[HttpServer]", r.RemoteAddr, message)
	Ligneous.Debug("[HttpServer]", "[Out] ", string(js))
	// set response headers
	w.Header().Set("Content-Type", "application/json")
	// allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// write response content
	w.Write(js)
}
