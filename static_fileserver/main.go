package main

import (
	"log"
	"net/http"
)

func main() {
  // by default, fileserver will look and serve index.html if one exists.
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
