package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetRandomQuote(quotesNumber int) []quote {
	var limit int
	var (
		output      []quote
		randomquote quote
	)

	if quotesNumber > len(Quotes) {
		limit = len(Quotes)
	} else {
		limit = quotesNumber
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < limit; i++ {
		randomquote = Quotes[rand.Intn(len(Quotes))]
		output = append(output, randomquote)
	}
	return output
}

func jsonQuotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	// Converting string to int so we can compare
	number, _ := strconv.Atoi(vars["quotes"])
	if number == 0 {
		number = 1
	}

	json, _ := json.Marshal(GetRandomQuote(number))
	fmt.Fprintf(w, string(json))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", jsonQuotesHandler)
	r.HandleFunc("/{quotes}", jsonQuotesHandler).Methods("GET")
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
