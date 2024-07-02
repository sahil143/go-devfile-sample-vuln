package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

var port = os.Getenv("PORT")

func main() {
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting Server")
	log.Printf("TEST_ENV_VAR : %s", os.Getenv("TEST_ENV_VAR"))
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path != "" {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	} else {
		fmt.Fprint(w, "Hello World!")
	}
}
