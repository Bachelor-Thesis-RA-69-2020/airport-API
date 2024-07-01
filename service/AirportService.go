package service

import (
	"airport-API/model"
	"airport-API/repository"
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
	if search == "" {
		return s.repo.FindAll()
	}

	search = strings.ToLower(search)
	airports, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []model.Airport
	for _, airport := range airports {
		if strings.Contains(strings.ToLower(airport.Name), search) || strings.Contains(strings.ToLower(airport.IATA), search) {
			result = append(result, airport)
		}
	}

	return result, nil
}
