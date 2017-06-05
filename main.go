package main

import "net/http"
import "fmt"

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello Bobbo")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8000", nil)  
}
