
async function getWeather() {

    // Get city name from input
    const city = document.getElementById("cityInput").value.trim();

    const message = document.getElementById("message");

    if (city === "") {
        message.innerText = "Please enter a city name";
        return;
    }

    message.innerText = "Loading weather data...";

    try {

        // Send request to Go backend
        const response = await fetch(
            `/weather?city=${encodeURIComponent(city)}`
        );

        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.error || "Unable to fetch weather");
        }

        // Display weather information
        document.getElementById("cityName").innerText =
            data.city + ", " + data.country;

        document.getElementById("temperature").innerText =
            data.weather.temperature_2m + " °C";

        document.getElementById("humidity").innerText =
            data.weather.relative_humidity_2m + "%";

        document.getElementById("wind").innerText =
            data.weather.wind_speed_10m + " km/h";

        document.getElementById("feelsLike").innerText =
            data.weather.apparent_temperature + " °C";

        // Convert weather code into a simple description
        const code = data.weather.weather_code;

        let condition = "Unknown";

        if (code === 0) {
            condition = "Clear Sky";
        } else if (code <= 3) {
            condition = "Partly Cloudy";
        } else if (code <= 48) {
            condition = "Foggy";
        } else if (code <= 67) {
            condition = "Rainy";
        } else if (code <= 77) {
            condition = "Snowy";
        } else if (code <= 82) {
            condition = "Rain Showers";
        } else if (code <= 86) {
            condition = "Snow Showers";
        } else {
            condition = "Thunderstorm";
        }

        document.getElementById("condition").innerText = condition;

        message.innerText = "";

    } catch (error) {

        message.innerText = error.message;

    }
}
