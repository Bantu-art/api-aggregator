package data

// FetchConcurrently fetches data from multiple URLs concurrently.
// It takes a map of URL identifiers and their corresponding URLs, 
// launches a goroutine for each URL, and collects the results in a 
// map with the URL identifiers as keys. If an error occurs while 
// fetching the data, the result for that URL will contain an error message.
//
// Arguments:
// - urls: A map where the keys are identifiers (e.g., names or types) 
//   and the values are URLs to fetch data from.
//
// Returns:
// - A map where the keys are the identifiers from the input map, 
//   and the values are the fetched data or error message.
func FetchConcurrently(urls map[string]string) map[string]string {
	results := make(map[string]string)
	ch := make(chan struct {
		url  string
		data string
	}, len(urls))

	// Launch a goroutine for each URL
	for key, url := range urls {
		go func(key, url string) {
			data, err := FetchData(url)
			if err != nil {
				ch <- struct {
					url  string
					data string
				}{key, "Error fetching data"}
			} else {
				ch <- struct {
					url  string
					data string
				}{key, data}
			}
		}(key, url)
	}

	// Collect responses from the channel
	for range urls {
		result := <-ch
		results[result.url] = result.data
	}

	return results
}
