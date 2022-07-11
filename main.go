package main

import (
  "github.com/joho/godotenv"
  "encoding/json"
  "net/http"
  "math/rand"
  "time"
  "fmt"
  "os"
)

func main(){
  var b []byte
  err := godotenv.Load()
  port := os.Getenv("PORT")

  if err != nil || port == "" {
    port = "8080"
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    b, _ = json.Marshal(Quotes[rand.Intn( len(Quotes) )])

    fmt.Fprintf(w, string(b))
  })

  fmt.Println("Server running on port", port)
	http.ListenAndServe(":" + port, nil)
}

