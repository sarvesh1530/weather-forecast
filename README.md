# 🌦️ Weather Forecast App

A simple and responsive Weather Forecast application built using **Go (Golang), Gin, HTML, CSS, and JavaScript**. It allows users to search for a city and view its current weather information using the Open-Meteo API.

## 🚀 Features

- Search weather by city name.
- Display current temperature.
- Show feels-like temperature.
- Display humidity.
- Show wind speed.
- Display weather conditions such as clear sky, rain, fog, and thunderstorms.
- Automatically detect the city's geographical coordinates.
- Handle invalid city names and API errors.

## 🛠️ Technologies Used

| Technology | Purpose |
|---|---|
| Go (Golang) | Backend development |
| Gin | Go web framework |
| HTML5 | Webpage structure |
| CSS3 | Styling and layout |
| JavaScript | Frontend functionality |
| Open-Meteo API | Weather and geocoding data |

## 📁 Project Structure

```text
weather-forecast/
│
├── backend/
│   ├── main.go
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── index.html
│   ├── style.css
│   └── script.js
│
├── .gitignore
└── README.md
```

## 🌐 APIs Used

This project uses the free Open-Meteo APIs.

**1. Geocoding API**

Converts the entered city name into geographical coordinates (latitude and longitude).

https://geocoding-api.open-meteo.com/v1/search

**2. Weather Forecast API**

Fetches current weather information using the geographical coordinates.

https://api.open-meteo.com/v1/forecast

No API key is required for the endpoints used in this project.

## ⚙️ Installation and Setup

### Prerequisites

- Go installed on your computer.
- Git installed on your computer.
- A web browser.

### Step 1: Clone the repository

```bash
git clone https://github.com/sarvesh1530/weather-forecast.git
```

### Step 2: Navigate to the project

```bash
cd weather-forecast
```

### Step 3: Navigate to the backend

```bash
cd backend
```

### Step 4: Download dependencies

```bash
go mod download
```

### Step 5: Run the application

```bash
go run .
```

### Step 6: Open the application

Open the following URL in your browser:

http://localhost:8080

Enter a city name and click the Search button to view its current weather.

## 🔄 How It Works

1. The user enters a city name in the frontend.
2. JavaScript sends a request to the Go backend.
3. The backend calls the Open-Meteo Geocoding API to find the city's coordinates.
4. The backend uses latitude and longitude to request current weather data.
5. The Go backend returns the weather information as JSON.
6. JavaScript displays the weather information on the webpage.

## 🔗 Backend API Endpoint

**Request:**

```http
GET /weather?city=Delhi
```

**Example response:**

```json
{
  "city": "Delhi",
  "country": "India",
  "weather": {
    "temperature_2m": 28.5,
    "apparent_temperature": 30.2,
    "relative_humidity_2m": 45,
    "wind_speed_10m": 10.2,
    "weather_code": 1
  }
}
```

*The values above are illustrative. Actual weather values depend on the API response.*

## 🎯 Learning Outcomes

Through this project, I practiced:

- Building a backend application using Go.
- Creating REST API endpoints using Gin.
- Connecting external APIs with a Go backend.
- Handling JSON responses.
- Integrating frontend JavaScript with a backend API.
- Managing project dependencies using Go modules.
- Using Git and GitHub for version control.

## 👨‍💻 Author

**Sarvesh Ranjan**

- GitHub: https://github.com/sarvesh1530
- LinkedIn: https://linkedin.com/in/sarvesh-ranjan-663131353

---

⭐ If you find this project useful, feel free to explore the repository!
