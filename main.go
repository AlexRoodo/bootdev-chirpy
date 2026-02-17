package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", appHandler())
	mux.HandleFunc("/healthz", healthHandler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

func healthHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "text/plain; charset=utf-8")
	resp.WriteHeader(200)
	resp.Write([]byte("OK"))
}

func appHandler() http.Handler {
	return http.StripPrefix("/app", http.FileServer(http.Dir("./")))
}
