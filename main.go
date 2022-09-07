package main

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

func jsonQuotesHandler(ctx *gin.Context) {
	param := ctx.Params.ByName("number")
	// converting from string to int
	number, _ := strconv.Atoi(param)
	if number <= 0 {
		number = 1
	}

	ctx.JSON(http.StatusOK, GetRandomQuote(number))

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET"},
		AllowAllOrigins: true,
		MaxAge:          20 * time.Minute,
	}))

	r.GET("/", jsonQuotesHandler)
	r.GET("/:number", jsonQuotesHandler)
	return r
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := setupRouter()
	r.Run(":" + port)
}
