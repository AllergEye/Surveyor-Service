package surveyor_dish

import (
	"context"

	"go.uber.org/zap"
)

type DishController interface {
	AddDishesToRestaurant(ctx context.Context, requestBody AddDishesToRestaurantRequestBody) error
}

type DishControllerImplementation struct {
	Logger      *zap.SugaredLogger
	DishService DishService
	Marshallers Marshallers
}

func NewDishContoller(logger *zap.SugaredLogger, dishService DishService, marshallers Marshallers) DishController {
	return DishControllerImplementation{
		Logger:      logger,
		DishService: dishService,
		Marshallers: marshallers,
	}
}

func (dc DishControllerImplementation) AddDishesToRestaurant(ctx context.Context, requestBody AddDishesToRestaurantRequestBody) error {
	restaurantId, dishes, err := dc.Marshallers.MarshalAddDishesToRestaurantRequestBody(requestBody)
	if err != nil {
		return err
	}

	err = dc.DishService.AddDishesToRestaurant(ctx, restaurantId, dishes)
	if err != nil {
		return err
	}
	return nil
}
