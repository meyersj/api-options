package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Success bool
	Msg     string
}

func build_response(r Response) (string, error) {
	// convert Response struct into json btyte array
	// then convert that as a json string and return it
	json_bytes, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	json_string := string(json_bytes[:])
	return json_string, nil
}

func authenticate(r *http.Request) (string, error) {
	var response Response
	params := r.URL.Query()
	user_id := params.Get("id")
	password := params.Get("pass")
	if user_id == "" || password == "" {
		response = Response{Success: false, Msg: "Missing id and/or pass parameters"}
	} else if user_id == "admin" && password == "admin" {
		response = Response{Success: true, Msg: "User authenticated"}
	} else {
		response = Response{Success: false, Msg: "Try user:admin password:admin"}
	}
	return build_response(response)
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	// set response headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "text/plain")
	// return text response
	fmt.Fprintf(w, "Index")
}

func user_handler(w http.ResponseWriter, r *http.Request) {
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

func start_server(port int) {
	port_str := ":" + strconv.Itoa(port)
	fmt.Println("Starting server at 127.0.0.1" + port_str)
	http.HandleFunc("/index", index_handler)
	http.HandleFunc("/user", user_handler)
	http.ListenAndServe(port_str, nil)
}

func main() {
	start_server(5001)
}
