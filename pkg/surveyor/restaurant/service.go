package surveyor

import (
	"errors"

	"github.com/allergeye/surveyor-service/internal/database"
	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	"go.uber.org/zap"
)

type RestaurantService interface {
	GetAllRestaurants() ([]restaurant.Restaurant, error)
	AddRestaurant(restaurant restaurant.Restaurant) error
}

type RestaurantServiceImplementation struct {
	logger         *zap.SugaredLogger
	restaurantRepo database.RestaurantRepository
}

var (
	ErrRestaurantAlreadyExists = errors.New("a restaurant with that name already exists")
)

func NewRestaurantService(logger *zap.SugaredLogger, restaurantRepo database.RestaurantRepository) RestaurantService {
	return RestaurantServiceImplementation{
		logger:         logger,
		restaurantRepo: restaurantRepo,
	}
}

func (s RestaurantServiceImplementation) GetAllRestaurants() ([]restaurant.Restaurant, error) {
	restaurants, err := s.restaurantRepo.GetAllRestaurants()
	if err != nil {
		return []restaurant.Restaurant{}, err
	}
	return restaurants, nil
}

func (s RestaurantServiceImplementation) AddRestaurant(restaurantToAdd restaurant.Restaurant) error {
	existingRestaurant, err := s.restaurantRepo.GetRestaurantByName(restaurantToAdd.Name)

	locationsToAdd := make([]restaurant.Location, 0)
	if err == nil {
		for _, location := range restaurantToAdd.Locations {
			found := false
			for _, existingLocation := range existingRestaurant.Locations {
				if s.locationsMatch(location, existingLocation) {
					found = true
					break
				}
			}
			if !found {
				locationsToAdd = append(locationsToAdd, location)
			}
		}

		if len(locationsToAdd) == 0 {
			return ErrRestaurantAlreadyExists
		}

		err = s.restaurantRepo.UpdateRestaurantLocations(*existingRestaurant, locationsToAdd)
		if err != nil {
			return err
		}
		return nil
	}

	err = s.restaurantRepo.AddRestaurant(restaurantToAdd)
	if err != nil {
		return err
	}

	return nil
}

func (s RestaurantServiceImplementation) locationsMatch(location restaurant.Location, targetLocation restaurant.Location) bool {
	return location.StreetAddressLine1 == targetLocation.StreetAddressLine1 && location.StreetAddressLine2 == targetLocation.StreetAddressLine2 && location.City == targetLocation.City && location.Province == targetLocation.Province && location.Country == targetLocation.Country && location.PostalCode == targetLocation.PostalCode
}
