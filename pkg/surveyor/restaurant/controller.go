package surveyor

import (
	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type RestaurantController interface {
	GetAllRestaurants() ([]restaurant.Restaurant, error)
	AddRestaurant(requestBody AddRestaurantRequestBody) error
}

type RestaurantControllerImplementation struct {
	logger            *zap.SugaredLogger
	restaurantService RestaurantService
}

func NewRestaurantController(logger *zap.SugaredLogger, restaurantService RestaurantService) RestaurantController {
	return RestaurantControllerImplementation{
		logger:            logger,
		restaurantService: restaurantService,
	}
}

func (c RestaurantControllerImplementation) GetAllRestaurants() ([]restaurant.Restaurant, error) {
	restaurants, err := c.restaurantService.GetAllRestaurants()
	if err != nil {
		return []restaurant.Restaurant{}, err
	}
	return restaurants, nil
}

func (c RestaurantControllerImplementation) AddRestaurant(requestBody AddRestaurantRequestBody) error {
	restaurant := restaurant.Restaurant{
		RestaurantId: uuid.New(),
		Name:         requestBody.Name,
		Locations:    requestBody.Locations,
	}

	err := c.restaurantService.AddRestaurant(restaurant)
	if err != nil {
		return err
	}

	return nil
}
