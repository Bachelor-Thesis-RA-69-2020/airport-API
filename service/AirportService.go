package service

import (
	"airport-API/model"
	"airport-API/repository"
	"sort"
	"strings"
)

type AirportService interface {
	GetAirports(search string) ([]model.Airport, error)
}

type airportService struct {
	repo repository.AirportRepository
}

func NewAirportService(repo repository.AirportRepository) AirportService {
	return &airportService{repo: repo}
}

func (s *airportService) GetAirports(search string) ([]model.Airport, error) {
	airports, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	if search == "" {
		sort.Slice(airports, func(i, j int) bool {
			return airports[i].IATA < airports[j].IATA
		})
		return airports, nil
	}

	search = strings.ToLower(search)
	var result []model.Airport
	for _, airport := range airports {
		if strings.Contains(strings.ToLower(airport.Name), search) || strings.Contains(strings.ToLower(airport.IATA), search) {
			result = append(result, airport)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].IATA < result[j].IATA
	})

	return result, nil
}
