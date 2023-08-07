package surveyor_dish

import (
	"context"

	"github.com/allergeye/surveyor-service/internal/database"
	"github.com/allergeye/surveyor-service/internal/domain/dish"
	"go.uber.org/zap"
)

type DishService interface {
	GetDishesByRestaurantId(ctx context.Context, restaurantId string) ([]dish.Dish, error)
	AddDishesToRestaurant(ctx context.Context, restaurantId string, dishes []dish.Dish) error
}

type DishServiceImplementation struct {
	Logger         *zap.SugaredLogger
	RestaurantRepo database.RestaurantRepository
	DishRepo       database.DishRepository
}

func NewDishService(logger *zap.SugaredLogger, restaurantRepo database.RestaurantRepository, dishRepo database.DishRepository) DishService {
	return DishServiceImplementation{
		Logger:         logger,
		RestaurantRepo: restaurantRepo,
		DishRepo:       dishRepo,
	}
}

func (ds DishServiceImplementation) GetDishesByRestaurantId(ctx context.Context, restaurantId string) ([]dish.Dish, error) {
	restaurant, err := ds.RestaurantRepo.GetRestaurantById(ctx, restaurantId)
	if err != nil {
		return []dish.Dish{}, err
	}
	dishes := make([]dish.Dish, len(restaurant.DishIds))
	for i, dishId := range restaurant.DishIds {
		dishToRetrieve, err := ds.DishRepo.GetDishById(ctx, dishId)
		if err != nil {
			return []dish.Dish{}, err
		}
		dishes[i] = *dishToRetrieve
	}
	return dishes, nil
}

func (ds DishServiceImplementation) AddDishesToRestaurant(ctx context.Context, restaurantId string, dishes []dish.Dish) error {
	dishIds := make([]string, len(dishes))
	for i, dish := range dishes {
		dishIds[i] = dish.DishId
	}

	err := ds.RestaurantRepo.AddDishesToRestaurant(ctx, restaurantId, dishIds)
	if err != nil {
		return err
	}

	err = ds.DishRepo.AddDishes(ctx, dishes)
	if err != nil {
		return err
	}

	return nil
}
