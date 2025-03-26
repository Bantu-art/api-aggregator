# API Aggregator

A Go-based API aggregator that fetches data from multiple sources concurrently, including news, stock market data, and weather updates. The project utilizes **Gin** for handling HTTP requests and **goroutines** for efficient concurrent data fetching.

## Features ✨  

- ✅ Fetches data from multiple APIs (News, Stocks, Weather).  
- ✅ Uses concurrent processing to improve response time.  
- ✅ Secure API key management using `.env` files.  
- ✅ RESTful API built with **Gin**.  

## Installation 🔧  

### 1️⃣ Clone the Repository  
```sh
git clone https://github.com/Bantu-art/api-aggregator.git
cd api-aggregator
```
# 2️⃣ Set Up Environment Variables
Create a .env file in the root directory and add your API keys:

```sh
NEWS_API_KEY=your_news_api_key
STOCK_API_KEY=your_stock_api_key
WEATHER_API_KEY=your_weather_api_key
```

# 3️⃣ Install Dependencies
```sh
go mod tidy
```

# 4️⃣ Run the Application
```sh
go run main.go
```

# API Endpoints 🌍
Fetch Aggregated Data
```sh
GET /data
```

## Response Format:

```sh
{
  "news": {...},
  "stocks": {...},
  "weather": {...}
}
```

# Project Structure 📂
```sh
api-aggregator/
│── data/              # Data fetching logic
│── main.go            # Entry point
│── .env               # API keys (not committed)
│── go.mod             # Go module dependencies
│── README.md          # Project documentation
```

# Future Improvements 🛠️
✅ Implement concurrent data fetching ➡️ DONE

⏳ Add caching for API responses

⏳ Improve error handling and logging

⏳ Add more data sources

Contributing 🤝
Feel free to fork, improve, and submit PRs!

License 📜
MIT License © 2025