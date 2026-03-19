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
	stripedPath := http.StripPrefix("/app", flPath)

	mux := http.NewServeMux()
	mux.Handle("/app/", stripedPath)

	mux.HandleFunc("/healthz", handleReadiness)

	svr := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Starting server on port: %s\n", port)
	log.Fatal(svr.ListenAndServe())

}

func handleReadiness(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
