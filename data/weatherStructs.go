package data

// WeatherResponse represents the JSON response from a weather API.
// It contains weather details, location coordinates, and atmospheric conditions.
type WeatherResponse struct {
	Coord      Coordinates `json:"coord"`       // Geographic coordinates (latitude and longitude)
	Weather    []Weather   `json:"weather"`     // Weather conditions (e.g., clear, cloudy, rainy)
	Base       string      `json:"base"`        // Internal API base station data
	Main       MainWeather `json:"main"`        // Main weather data such as temperature and humidity
	Visibility int         `json:"visibility"`  // Visibility in meters
	Wind       Wind        `json:"wind"`        // Wind speed, direction, and gusts
	Rain       Rain        `json:"rain,omitempty"` // Rain data (optional, present only if rain is recorded)
	Clouds     Clouds      `json:"clouds"`      // Cloud cover percentage
	Dt         int64       `json:"dt"`          // Time of data calculation (Unix timestamp)
	Sys        SystemInfo  `json:"sys"`         // Additional system information such as country and sunrise/sunset times
	Timezone   int         `json:"timezone"`    // Timezone offset in seconds from UTC
	ID         int         `json:"id"`          // City ID
	Name       string      `json:"name"`        // City name
	Cod        int         `json:"cod"`         // API response code (e.g., 200 for success)
}

// Coordinates represents the latitude and longitude of a location.
type Coordinates struct {
	Lon float64 `json:"lon"` // Longitude coordinate
	Lat float64 `json:"lat"` // Latitude coordinate
}

// Weather represents the current weather conditions.
type Weather struct {
	ID          int    `json:"id"`          // Weather condition ID
	Main        string `json:"main"`        // Main weather description (e.g., Rain, Snow, Clear)
	Description string `json:"description"` // Detailed weather description
	Icon        string `json:"icon"`        // Weather icon ID for graphical representation
}

// MainWeather contains temperature and atmospheric data.
type MainWeather struct {
	Temp        float64 `json:"temp"`        // Current temperature in Kelvin
	FeelsLike   float64 `json:"feels_like"`  // Perceived temperature
	TempMin     float64 `json:"temp_min"`    // Minimum temperature
	TempMax     float64 `json:"temp_max"`    // Maximum temperature
	Pressure    int     `json:"pressure"`    // Atmospheric pressure in hPa
	Humidity    int     `json:"humidity"`    // Humidity percentage
	SeaLevel    int     `json:"sea_level"`   // Sea level atmospheric pressure (optional)
	GroundLevel int     `json:"grnd_level"`  // Ground level atmospheric pressure (optional)
}

// Wind represents wind speed, direction, and gusts.
type Wind struct {
	Speed float64 `json:"speed"` // Wind speed in meters per second
	Deg   int     `json:"deg"`   // Wind direction in degrees
	Gust  float64 `json:"gust"`  // Wind gust speed (optional)
}

// Rain represents precipitation levels.
type Rain struct {
	OneHour float64 `json:"1h"` // Rainfall volume for the last 1 hour in mm
}

// Clouds represents cloud coverage data.
type Clouds struct {
	All int `json:"all"` // Cloudiness percentage
}

// SystemInfo contains additional metadata such as country code and sunrise/sunset times.
type SystemInfo struct {
	Country string `json:"country"` // Country code (e.g., "US" for the United States)
	Sunrise int64  `json:"sunrise"` // Sunrise time (Unix timestamp)
	Sunset  int64  `json:"sunset"`  // Sunset time (Unix timestamp)
}
