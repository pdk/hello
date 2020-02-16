package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	log.Printf("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// HelloServer responds with a friendly greeting.
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Saluton mondo!!")
}
