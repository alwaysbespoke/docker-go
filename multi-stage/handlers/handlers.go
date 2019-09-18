package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// EchoData: echoed request data
type EchoData struct {
	Headers http.Header `json:"request_headers"`
	Body    string      `json:"request_body"`
	Method  string      `json:"request_method"`
	Path    string      `json:"request_path"`
}

// EchoHandler returns various data from the incoming request
func EchoHandler(response http.ResponseWriter, request *http.Request) {

	fmt.Println(os.Getenv("PORT"))
	fmt.Println(os.Getenv("HOST"))
	fmt.Println("REQUEST:\n", request)
	fmt.Println()

	// only allow http 1.1
	if request.Proto != "HTTP/1.1" {
		fmt.Println("Error: invalid request protocol")
		response.WriteHeader(400)
		return
	}

	// only allow Get and Post
	if request.Method != http.MethodGet && request.Method != http.MethodPost {
		fmt.Println("Error: invalid request method")
		response.WriteHeader(400)
		return
	}

	// read body
	var body []byte
	var err error
	if request.Method == http.MethodPost {
		// TODO -> check content type in header
		// TODO -> validate JSON
		body, err = ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			response.WriteHeader(500)
			return
		}
	}

	// populate struct with request data
	var echoData EchoData
	echoData.Headers = request.Header
	echoData.Method = request.Method
	echoData.Path = request.URL.Path
	echoData.Body = string(body)

	// marshal struct into json
	j, err := json.Marshal(echoData)
	if err != nil {
		fmt.Println("Error: ", err)
		response.WriteHeader(500)
		return
	}

	// write response
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	response.Write(j)

}
