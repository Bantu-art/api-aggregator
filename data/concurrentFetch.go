package data

import (
	"encoding/json"
	"time"
)

// APIResponse represents the combined response from all APIs
type APIResponse struct {
	Weather WeatherResponse `json:"weather"`
	News    NewsResponse    `json:"news"`
	Stock   StockResponse   `json:"stock"`
}

// FetchConcurrently fetches data from multiple URLs concurrently and caches the results.
// It takes a map of URL identifiers and their corresponding URLs,
// launches a goroutine for each URL, and collects the results in a
// map with the URL identifiers as keys. If an error occurs while
// fetching the data, the result for that URL will contain an error message.
//
// Arguments:
//   - urls: A map where the keys are identifiers (e.g., names or types)
//     and the values are URLs to fetch data from.
//
// Returns:
//   - A map where the keys are the identifiers from the input map,
//     and the values are the fetched data or error message.
func FetchConcurrently(urls map[string]string) (*APIResponse, error) {
	// Try to get cached data first
	var cachedResponse APIResponse
	err := GetCachedData("combined_data", &cachedResponse)
	if err == nil {
		return &cachedResponse, nil
	}

	// If no cached data, fetch from APIs
	results := make(map[string]interface{})
	ch := make(chan struct {
		key  string
		data interface{}
		err  error
	}, len(urls))

	// Launch a goroutine for each URL
	for key, url := range urls {
		go func(key, url string) {
			var result interface{}
			var err error

			switch key {
			case "weather":
				var weather WeatherResponse
				err = fetchAndParse(url, &weather)
				result = weather
			case "news":
				var news NewsResponse
				err = fetchAndParse(url, &news)
				// Limit to 5 articles
				if len(news.Articles) > 5 {
					news.Articles = news.Articles[:5]
				}
				result = news
			case "stock":
				var stock StockResponse
				err = fetchAndParse(url, &stock)
				// Limit to 5 stock entries
				if len(stock.Data) > 5 {
					stock.Data = stock.Data[:5]
				}
				result = stock
			}

			ch <- struct {
				key  string
				data interface{}
				err  error
			}{key, result, err}
		}(key, url)
	}

	// Collect responses from the channel
	for range urls {
		result := <-ch
		if result.err != nil {
			return nil, result.err
		}
		results[result.key] = result.data
	}

	// Combine results into APIResponse
	response := &APIResponse{
		Weather: results["weather"].(WeatherResponse),
		News:    results["news"].(NewsResponse),
		Stock:   results["stock"].(StockResponse),
	}

	// Cache the results for 5 minutes
	err = SetCachedData("combined_data", response, 5*time.Minute)
	if err != nil {
		// Log the error but don't fail the request
		// TODO: Add proper logging
	}

	return response, nil
}

// fetchAndParse fetches data from a URL and unmarshals it into the provided interface
func fetchAndParse(url string, result interface{}) error {
	data, err := FetchData(url)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), result)
}
