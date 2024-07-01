package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"airport-API/model"
)

type AirportRepository interface {
	FindAll() ([]model.Airport, error)
}

type airportRepository struct {
	dataFile string
}

func NewAirportRepository(dataFile string) AirportRepository {
	return &airportRepository{dataFile: dataFile}
}

func (r *airportRepository) loadAirports() ([]model.Airport, error) {
	file, err := os.Open(r.dataFile)
	if err != nil {
		return nil, fmt.Errorf("error while opening file '%s': %v", r.dataFile, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error while reading CSV file: %v", err)
	}

	var airports []model.Airport
	for idx, record := range records {
		if idx == 0 {
			continue
		}

		if isEmpty(record[2]) || isEmpty(record[3]) || isEmpty(record[13]) || isEmpty(record[4]) || isEmpty(record[5]) || isEmpty(record[6]) || isEmpty(record[7]) || isEmpty(record[8]) || isEmpty(record[9]) || isEmpty(record[10]) {
			continue
		}

		if record[2] != "large_airport" {
			continue
		}

		airport := model.Airport{
			Name:         record[3],
			IATA:         record[13],
			Latitude:     record[4],
			Longitude:    record[5],
			Elevation:    record[6],
			Continent:    record[7],
			Country:      record[8],
			Region:       record[9],
			Municipality: record[10],
		}

		airports = append(airports, airport)
	}

	return airports, nil
}

func (r *airportRepository) FindAll() ([]model.Airport, error) {
	return r.loadAirports()
}

func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
