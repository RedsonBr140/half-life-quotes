package main

import (
"encoding/json"
"fmt"
"math/rand"
"net/http"
"os"
"time"
)

func main(){
  var(
    port string = os.Getenv("PORT")
    b    []byte
  )

  if port == "" {
    port = "8080"
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    rand.Seed(time.Now().UnixNano())
    b, _ = json.Marshal(Quotes[rand.Intn( len(Quotes) )])
    fmt.Fprintf(w, string(b))
  })

  fmt.Println("Server running on port", port)
  http.ListenAndServe(":" + port, nil)
}

