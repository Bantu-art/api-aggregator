// Package data provides structures to represent news data responses from an API.
package data

// NewsResponse represents the response received from a news API.
// It contains the status of the request, the total number of results, and a list of articles.
type NewsResponse struct {
	Status       string    `json:"status"`       // Status of the API response (e.g., "ok" or "error").
	TotalResults int       `json:"totalResults"` // The total number of news articles returned by the API.
	Articles     []Article `json:"articles"`     // A list of news articles retrieved from the API.
}

// Article represents a single news article retrieved from the API.
// It contains details about the article, including its source, author, title, description, and content.
type Article struct {
	Source      Source `json:"source"`      // The source of the news article.
	Author      string `json:"author"`      // The author of the article (if available).
	Title       string `json:"title"`       // The title of the news article.
	Description string `json:"description"` // A brief description or summary of the article.
	URL         string `json:"url"`         // The URL of the full article.
	URLToImage  string `json:"urlToImage"`  // The URL of the article's main image (if available).
	PublishedAt string `json:"publishedAt"` // The date and time when the article was published.
	Content     string `json:"content"`     // The full content of the article (if available).
}

// Source represents the source of a news article, including its ID and name.
type Source struct {
	ID   string `json:"id"`   // The unique identifier for the news source (if available).
	Name string `json:"name"` // The name of the news source.
}
