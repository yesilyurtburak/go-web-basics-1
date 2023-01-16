package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = "8080"
const ipAddress = "127.0.0.1"

var url = fmt.Sprintf("%s:%s", ipAddress, portNumber)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	nb, err := fmt.Fprintf(w, "Hello Web")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of bytes written: %d\n", nb)
}

func main() {
	// routing
	http.HandleFunc("/", HomeHandler)

	// listening http requests
	fmt.Printf("Listening traffic at %s\n", url)
	err := http.ListenAndServe(url, nil)
	log.Fatal(err)
}
