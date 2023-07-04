package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Manta Ray is listening on port 3000")
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
