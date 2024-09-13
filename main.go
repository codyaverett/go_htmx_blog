package main

import (
  "fmt"
  "log"
  "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World")
}

func main() {
  // Define routes
  http.HandleFunc("/", helloHandler)

  // Start the server
  fmt.Println("Starting server at port 8080...")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}

