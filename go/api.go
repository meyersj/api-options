package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// http://www.freshblurbs.com/blog/2013/12/07/hello-web-golang.html

type Response struct {
	Success bool
	Msg     string
}

func authenticate(r *http.Request) (string, error) {
	var response Response
	params := r.URL.Query()
	id := params.Get("id")
	pass := params.Get("pass")
	msg := "User authenticated"
	if id == "admin" && pass == "admin" {
		response = Response{Success: true, Msg: msg}
	} else {
		msg = "Try user:admin password:admin"
		response = Response{Success: false, Msg: msg}
	}
	return json_response(response)
}

func view_handler(w http.ResponseWriter, r *http.Request) {
	// set response headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")
	// authenticate user and build json response
	response, err := authenticate(r)
	if err != nil {
		http.Error(w, "Oops", http.StatusInternalServerError)
	}
	// send response as json
	fmt.Fprintf(w, response)
}

func json_response(r Response) (string, error) {
	// convert Response struct into json btyte array
	// then convert that as a json string and return it
	json_bytes, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	json_string := string(json_bytes[:])
	return json_string, nil
}

func start_server(port int) {
	port_str := ":" + strconv.Itoa(port)
	fmt.Println("Starting server at 127.0.0.1" + port_str)
	http.HandleFunc("/user", view_handler)
	http.ListenAndServe(port_str, nil)
}

func main() {
	start_server(8001)
}
