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

func (s RestaurantServiceImplementation) AddRestaurant(restaurant restaurant.Restaurant) error {
	_, err := s.restaurantRepo.GetRestaurantByName(restaurant.Name)
	if err == nil {
		return ErrRestaurantAlreadyExists
	}

	err = s.restaurantRepo.AddRestaurant(restaurant)
	if err != nil {
		return err
	}

	return nil
}
