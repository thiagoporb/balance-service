package main

import (
	"fmt"
	"log"
	"net/http"
)

// request handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/api/balances", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
