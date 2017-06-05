package main

import "net/http"
import "fmt"

func handler(w http.Responsewriter, r *http.Request) {
  fmt.Fprintf(w, "Hello Bobbo")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":80", nil)  
}
