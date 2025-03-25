package data

import (
	"fmt"
	"io"
	"net/http"
)

// FetchData retrieves data from the given URL using an HTTP GET request.
//
// Parameters:
//   - url: A string representing the API endpoint or web resource to fetch data from.
//
// Returns:
//   - A string containing the response body if the request is successful.
//   - An error if the request fails or the response body cannot be read.
//
// Example Usage:
//
//	response, err := FetchData("https://api.example.com/data")
//	if err != nil {
//	    log.Println("Error fetching data:", err)
//	} else {
//	    fmt.Println("Fetched Data:", response)
//	}
func FetchData(url string) (string, error) {
	// Send HTTP GET request to the specified URL
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}
	// Ensure the response body is closed after function execution
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Return the response as a string
	return string(body), nil
}
