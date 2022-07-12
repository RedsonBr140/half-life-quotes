package main

import (
"encoding/json"
"fmt"
"math/rand"
"net/http"
"os"
"time"
)

var b []byte

func Handler(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Access-Control-Allow-Origin", "*")
  rand.Seed(time.Now().UnixNano())

  b, _ = json.Marshal(Quotes[rand.Intn( len(Quotes) )])
  fmt.Fprintf(w, string(b))
}

func main(){
  port := os.Getenv("PORT")

  if port == "" {
    port = "8080"
  }
  http.HandleFunc("/", Handler)
  fmt.Println("Server running on port", port)
  http.ListenAndServe(":" + port, nil)
}

