package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	flPath := http.FileServer(http.Dir(filepathRoot))

	mux := http.NewServeMux()
	mux.Handle("/", flPath)

	svr := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Starting server on port: %s\n", port)
	log.Fatal(svr.ListenAndServe())

}
