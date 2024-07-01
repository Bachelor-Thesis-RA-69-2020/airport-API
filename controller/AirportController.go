package controller

import (
	"airport-API/service"
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AirportController interface {
	GetAirports(c *gin.Context)
}

type airportController struct {
	service service.AirportService
}

func NewAirportController(svc service.AirportService) AirportController {
	return &airportController{
		service: svc,
	}
}

func (ac *airportController) GetAirports(c *gin.Context) {
	search := c.Query("search")
	airports, err := ac.service.GetAirports(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch airports"})
		return
	}

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", `attachment; filename="airports.csv"`)

	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	header := []string{"Name", "IATA", "Latitude", "Longitude", "Elevation", "Continent", "Country", "Region", "Municipality"}
	writer.Write(header)

	for _, airport := range airports {
		row := []string{
			airport.Name,
			airport.IATA,
			airport.Latitude,
			airport.Longitude,
			airport.Elevation,
			airport.Continent,
			airport.Country,
			airport.Region,
			airport.Municipality,
		}
		writer.Write(row)
	}
}
