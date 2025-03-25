package data

// StockResponse represents the API response for stock market data.
// It contains pagination details and a list of stock data records.
type StockResponse struct {
	Pagination Pagination  `json:"pagination"` // Metadata for paginating the stock data
	Data       []StockData `json:"data"`       // List of stock market data
}

// Pagination provides details about the paginated response.
type Pagination struct {
	Limit  int `json:"limit"`  // Maximum number of records per page
	Offset int `json:"offset"` // Position of the first record in the current page
	Count  int `json:"count"`  // Number of records returned in this response
	Total  int `json:"total"`  // Total number of available records
}

// StockData represents a single stock market data entry.
type StockData struct {
	Open          float64 `json:"open"`           // Opening price of the stock
	High          float64 `json:"high"`           // Highest price of the stock during the day
	Low           float64 `json:"low"`            // Lowest price of the stock during the day
	Close         float64 `json:"close"`          // Closing price of the stock
	Volume        float64 `json:"volume"`         // Number of shares traded
	AdjHigh       float64 `json:"adj_high"`       // Adjusted highest price
	AdjLow        float64 `json:"adj_low"`        // Adjusted lowest price
	AdjClose      float64 `json:"adj_close"`      // Adjusted closing price
	AdjOpen       float64 `json:"adj_open"`       // Adjusted opening price
	AdjVolume     float64 `json:"adj_volume"`     // Adjusted volume of shares traded
	SplitFactor   float64 `json:"split_factor"`   // Stock split factor
	Dividend      float64 `json:"dividend"`       // Dividend payout
	Name          string  `json:"name"`           // Company name
	ExchangeCode  string  `json:"exchange_code"`  // Exchange code where the stock is listed
	AssetType     string  `json:"asset_type"`     // Type of financial asset (e.g., stock, ETF)
	PriceCurrency string  `json:"price_currency"` // Currency in which the stock is priced
	Symbol        string  `json:"symbol"`         // Ticker symbol of the stock
	Exchange      string  `json:"exchange"`       // Name of the stock exchange
	Date          string  `json:"date"`           // Date of the stock data
}
