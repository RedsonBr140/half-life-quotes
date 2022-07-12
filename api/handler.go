package handler

import (
  "encoding/json"
  "net/http"
  "math/rand"
  "time"
  "fmt"
)
var b []byte

func Handler(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Access-Control-Allow-Origin", "*")
  rand.Seed(time.Now().UnixNano())

  b, _ = json.Marshal(Quotes[rand.Intn( len(Quotes) )])
  fmt.Fprintf(w, string(b))
}
