package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", status)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func status(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
