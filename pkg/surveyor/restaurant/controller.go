package surveyor

import (
	"context"
	"errors"

	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	"go.uber.org/zap"
)

var (
	ErrInvalidAllergen = errors.New("invalid allergen")
)

type RestaurantController interface {
	GetAllRestaurants(ctx context.Context) ([]restaurant.Restaurant, error)
	AddRestaurant(ctx context.Context, requestBody AddRestaurantRequestBody) error
}

type RestaurantControllerImplementation struct {
	Logger            *zap.SugaredLogger
	RestaurantService RestaurantService
	Marshallers       Marshallers
}

func NewRestaurantController(logger *zap.SugaredLogger, restaurantService RestaurantService, marshallers Marshallers) RestaurantController {
	return RestaurantControllerImplementation{
		Logger:            logger,
		RestaurantService: restaurantService,
		Marshallers:       marshallers,
	}
}

func (c RestaurantControllerImplementation) GetAllRestaurants(ctx context.Context) ([]restaurant.Restaurant, error) {
	restaurants, err := c.RestaurantService.GetAllRestaurants(ctx)
	if err != nil {
		return []restaurant.Restaurant{}, err
	}
	return restaurants, nil
}

func (c RestaurantControllerImplementation) AddRestaurant(ctx context.Context, requestBody AddRestaurantRequestBody) error {
	restaurant, dishes, err := c.Marshallers.MarshalRestaurantRequestBody(requestBody)
	if err != nil {
		return err
	}

	err = c.RestaurantService.AddRestaurant(ctx, *restaurant, dishes)
	if err != nil {
		return err
	}

	return nil
}
