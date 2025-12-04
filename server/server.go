package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloWorlResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8000
	http.HandleFunc("/hello-world", helloWorldHandler)

	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloWorlResponse{Message: "Hello World!"}

	data, err := json.Marshal(response)
	if err != nil {
		panic("Oops")
	}

	fmt.Fprint(w, string(data))
}
