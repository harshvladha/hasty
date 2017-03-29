package main

import (
	"encoding/json"
	"fmt"
	"github.com/hasty"
	"net/http"
)

func main() {
	fmt.Println("Starting...")
	myMux := hasty.New()
	myMux.Get("/test", http.HandlerFunc(getHandler))
	myMux.Post("/test", http.HandlerFunc(postHandler))
	fmt.Println("Registered...")
	err := myMux.ListenAndServe("8080")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func getHandler(rw http.ResponseWriter, req *http.Request) {
	var queryId string
	if query := req.URL.Query()["id"]; query != nil {
		queryId = query[0]
	}
	var myJson = struct {
		First  string
		Second string
		Third  string
	}{"Hello", "World!", queryId}

	json.NewEncoder(rw).Encode(&myJson)
}


func postHandler(rw http.ResponseWriter, req *http.Request) {

	var myJson = struct {
		First  string
		Second string
	}{"Hello", "Post World!"}

	json.NewEncoder(rw).Encode(&myJson)
}