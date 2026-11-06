
package main

import (
	"strconv"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

// Location response from Geocoding API
type Location struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeoResponse struct {
	Results []Location `json:"results"`
}

// Weather response from Forecast API
type WeatherResponse struct {
	Current struct {
		Time          string  `json:"time"`
		Temperature   float64 `json:"temperature_2m"`
		FeelsLike     float64 `json:"apparent_temperature"`
		Humidity      int     `json:"relative_humidity_2m"`
		WindSpeed     float64 `json:"wind_speed_10m"`
		WeatherCode   int     `json:"weather_code"`
	} `json:"current"`
}

func getWeather(c *gin.Context) {

	// Step 1: Get city name
	city := c.Query("city")

	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please enter a city name",
		})
		return
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	// Step 2: Find city coordinates
	geoURL := "https://geocoding-api.open-meteo.com/v1/search"

	geoParams := url.Values{}
	geoParams.Add("name", city)
	geoParams.Add("count", "1")
	geoParams.Add("language", "en")

	geoRequest, err := http.NewRequestWithContext(
		c.Request.Context(),
		http.MethodGet,
		geoURL+"?"+geoParams.Encode(),
		nil,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create location request",
		})
		return
	}

	geoResponse, err := client.Do(geoRequest)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Unable to connect to location service",
		})
		return
	}
	defer geoResponse.Body.Close()

	if geoResponse.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Location service returned an error",
		})
		return
	}

	var locationData GeoResponse

	err = json.NewDecoder(geoResponse.Body).Decode(&locationData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not read location data",
		})
		return
	}

	if len(locationData.Results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "City not found",
		})
		return
	}

	location := locationData.Results[0]

	// Step 3: Fetch weather using coordinates
	weatherURL := "https://api.open-meteo.com/v1/forecast"

	weatherParams := url.Values{}
	

	// Add latitude and longitude
	weatherParams.Set("latitude", formatCoordinate(location.Latitude))
	weatherParams.Set("longitude", formatCoordinate(location.Longitude))

	weatherParams.Add(
		"current",
		"temperature_2m,apparent_temperature,relative_humidity_2m,wind_speed_10m,weather_code",
	)

	weatherParams.Add("timezone", "auto")

	weatherRequest, err := http.NewRequestWithContext(
		c.Request.Context(),
		http.MethodGet,
		weatherURL+"?"+weatherParams.Encode(),
		nil,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create weather request",
		})
		return
	}

	weatherResponse, err := client.Do(weatherRequest)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Unable to connect to weather service",
		})
		return
	}
	defer weatherResponse.Body.Close()

	if weatherResponse.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Weather service returned an error",
		})
		return
	}

	// Step 4: Read weather JSON
	var weatherData WeatherResponse

	err = json.NewDecoder(weatherResponse.Body).Decode(&weatherData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not read weather data",
		})
		return
	}

	// Step 5: Return final result
	c.JSON(http.StatusOK, gin.H{
		"city":    location.Name,
		"country": location.Country,
		"weather": weatherData.Current,
	})
}

// Convert coordinates into strings
func formatCoordinate(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func main() {

	router := gin.Default()

	router.GET("/weather", getWeather)
	router.Static("/static", "../frontend")
router.StaticFile("/", "../frontend/index.html")

	router.Run(":8080")
}
