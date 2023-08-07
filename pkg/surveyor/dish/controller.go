package surveyor_dish

import (
	"context"
	"errors"

	"github.com/allergeye/surveyor-service/internal/domain/dish"
	"go.uber.org/zap"
)

var (
	ErrInvalidRestaurantId = errors.New("invalid restaurant id")
)

type DishController interface {
	GetDishesByRestaurantId(ctx context.Context, restaurantId string) ([]dish.Dish, error)
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

func (dc DishControllerImplementation) GetDishesByRestaurantId(ctx context.Context, restaurantId string) ([]dish.Dish, error) {
	dishes, err := dc.DishService.GetDishesByRestaurantId(ctx, restaurantId)
	if err != nil {
		return []dish.Dish{}, err
	}
	return dishes, nil
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
