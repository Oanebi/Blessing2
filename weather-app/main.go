package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// 1. Define the structural "shape" of our data
type WeatherRequest struct {
	City string `json:"city"`
}

type WeatherResponse struct {
	City             string  `json:"city"`
	CurrentTemp      float64 `json:"current_temperature"`
	PredictedTempTom float64 `json:"predicted_temperature_tomorrow"`
	Analysis         string  `json:"analysis"`
}

func main() {
	app := fiber.New()

	// 2. Enable CORS so your JavaScript frontend can talk to this backend
	app.Use(cors.New())

	// 3. Define the POST endpoint for our frontend to hit
	app.Post("/api/predict-weather", func(c *fiber.Ctx) error {
		
		// 4. Parse the incoming request from JavaScript
		req := new(WeatherRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// 5. THE BACKEND LOGIC (The "Kitchen")
		// Simulate fetching current weather data (e.g., 22.5°C)
		rand.Seed(time.Now().UnixNano())
		currentTemp := 15.0 + rand.Float64()*(35.0-15.0) 

		// 6. THE ALGORITHM / ML LOGIC
		// A simple predictive algorithm: if it's hot, tomorrow will likely cool down slightly
		var predictedTemp float64
		var analysis string

		if currentTemp > 28.0 {
			predictedTemp = currentTemp - (rand.Float64() * 3.0)
			analysis = "Heatwave detected. Predictive model expects a slight cooling trend tomorrow."
		} else {
			predictedTemp = currentTemp + (rand.Float64() * 2.5)
			analysis = "Stable conditions. Model predicts minor warming tomorrow."
		}

		// 7. Prepare the final structured package
		response := WeatherResponse{
			City:             req.City,
			CurrentTemp:      round(currentTemp, 1),
			PredictedTempTom: round(predictedTemp, 1),
			Analysis:         analysis,
		}

		// 8. Send it back to JavaScript as JSON
		return c.JSON(response)
	})

	// Start the server on port 3000
	app.Listen(":3000")
}

// Helper function to keep numbers clean (e.g., 24.5 instead of 24.5342342)
func round(val float64, precision int) float64 {
	return float64(int(val*10)) / 10
}