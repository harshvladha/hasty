package main

import (
	"encoding/json"
	"fmt"
	"github.com/hasty"
	"net/http"
)

func main() {
	myMux := hasty.New()
	myMux.Get("/test", http.HandlerFunc(get2Handler))
	myMux.Get("/test/:var1", http.HandlerFunc(get2Handler))
	myMux.Get("/test/:var1/info", http.HandlerFunc(get2InfoHandler))
	myMux.GetFunc("/home/:var2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Value of var2: "+hasty.GetValue(r, "var2"))
	})
	fmt.Println("Registered...")
	err := myMux.ListenAndServe("8080")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func get2Handler(rw http.ResponseWriter, req *http.Request) {
	var queryId string
	if query := hasty.GetValue(req, "var1"); query != "" {
		queryId = query
	}
	var myJson = struct {
		First  string
		Second string
		Third  string
	}{"Hello", "World!", queryId}

	json.NewEncoder(rw).Encode(&myJson)
}

func get2InfoHandler(rw http.ResponseWriter, req *http.Request) {
	var queryId string
	if query := hasty.GetValue(req, "var1"); query != "" {
		queryId = query
	}
	var myJson = struct {
		First string
		Third string
	}{"Info", queryId}

	json.NewEncoder(rw).Encode(&myJson)
}
