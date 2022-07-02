package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	port := "8080"
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// read html file as string
	bytes, _ := os.ReadFile("index.html")
	_, err := fmt.Fprint(w, string(bytes))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
