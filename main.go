package main

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"time"
	"fmt"

)

func main(){
  var b []byte
	
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    b, _ = json.Marshal(Quotes[rand.Intn( len(Quotes) )])
    fmt.Fprintf(w, string(b))
  })

	http.ListenAndServe(":8080", nil)
}

