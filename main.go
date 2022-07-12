package main

import (
  "fmt"
  "net/http"
  "os"
  "hl-api/api"
)

func main(){
  port := os.Getenv("PORT")

  if port == "" {
    port = "8080"
  }
  http.HandleFunc("/", handler.Handler)
  fmt.Println("Server running on port", port)
  http.ListenAndServe(":" + port, nil)
}

