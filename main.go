package main

import (
	"airport-API/controller"
	"airport-API/repository"
	"airport-API/service"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(currentFile))

	// Construct the absolute path to airports.csv
	absPath := filepath.Join(projectRoot, "airport-API", "repository", "airports.csv")
	fmt.Println("ALOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Println(absPath)

	r := gin.Default()

	airportRepo := repository.NewAirportRepository(absPath)
	airportService := service.NewAirportService(airportRepo)
	airportController := controller.NewAirportController(airportService)

	r.GET("/airports", airportController.GetAirports)

	port := ":8084"
	log.Printf("Starting server on port %s...\n", port)
	err := r.Run(port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
