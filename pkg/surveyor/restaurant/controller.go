package surveyor

import (
	"context"

	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	"github.com/allergeye/surveyor-service/internal/lib"
	"go.uber.org/zap"
)

type RestaurantController interface {
	GetAllRestaurants(ctx context.Context) ([]restaurant.Restaurant, error)
	AddRestaurant(ctx context.Context, requestBody AddRestaurantRequestBody) error
}

type RestaurantControllerImplementation struct {
	Logger            *zap.SugaredLogger
	RestaurantService RestaurantService
	Helpers           lib.Helpers
}

func NewRestaurantController(logger *zap.SugaredLogger, restaurantService RestaurantService, helpers lib.Helpers) RestaurantController {
	return RestaurantControllerImplementation{
		Logger:            logger,
		RestaurantService: restaurantService,
		Helpers:           helpers,
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
	restaurant := restaurant.Restaurant{
		RestaurantId: c.Helpers.GenerateUUID(),
		Name:         requestBody.Name,
		Locations:    requestBody.Locations,
	}

	err := c.RestaurantService.AddRestaurant(ctx, restaurant)
	if err != nil {
		return err
	}

	return nil
}
