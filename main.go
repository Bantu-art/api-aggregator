package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Bantu-art/api-aggregator/data"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Package main initializes and runs a web server that fetches data 
// concurrently from multiple external APIs, including weather, news, and stock market data.
//
// The server exposes a `/data` endpoint, which retrieves JSON data from 
// OpenWeatherMap, NewsAPI, and MarketStack APIs. The fetched data is
// returned as a JSON response.
//
// Environment Variables:
// - WEATHER_API_KEY: API key for OpenWeatherMap
// - NEWS_API_KEY: API key for NewsAPI
// - STOCK_API_KEY: API key for MarketStack
//
// The server listens on port 8080.
func main() {
	// Load environment variables once at startup
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	weatherAPIKey := os.Getenv("WEATHER_API_KEY")
	newsAPIKey := os.Getenv("NEWS_API_KEY")
	stockAPIKey := os.Getenv("STOCK_API_KEY")

	if weatherAPIKey == "" || newsAPIKey == "" || stockAPIKey == "" {
		log.Fatal("Missing one or more API keys")
	}

	urls := map[string]string{
		"weather": "https://api.openweathermap.org/data/2.5/weather?q=Nairobi&appid=" + weatherAPIKey,
		"news":    "https://newsapi.org/v2/top-headlines?country=us&apiKey=" + newsAPIKey,
		"stock":   "http://api.marketstack.com/v2/eod?symbols=AAPL&date_from=2025-03-15&date_to=2025-03-25&access_key=" + stockAPIKey,
	}
	r := gin.Default()

	r.GET("/data", func(c *gin.Context) {
		// Fetch the JSON data
		newsData := data.FetchConcurrently(urls)

		// Send the fetched data as a JSON response
		c.JSON(http.StatusOK, newsData)
	})

	r.Run(":8080") 
}
