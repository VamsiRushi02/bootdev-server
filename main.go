package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()

	svr := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Starting server on port: %s\n", port)
	log.Fatal(svr.ListenAndServe())

}
