package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("ok"))
  })
  mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
    _, _ = w.Write([]byte("hello from ploy"))
  })
  port := os.Getenv("PORT")
  if port == "" { port = "8080" }
  addr := ":" + port
  log.Printf("listening on %s", addr)
  log.Fatal(http.ListenAndServe(addr, mux))
}
