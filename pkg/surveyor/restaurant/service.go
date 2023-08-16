package surveyor_restaurant

import (
	"context"
	"errors"

	"github.com/allergeye/surveyor-service/internal/database"
	"github.com/allergeye/surveyor-service/internal/domain/dish"
	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	"go.uber.org/zap"
)

type RestaurantService interface {
	GetAllRestaurants(ctx context.Context) ([]restaurant.Restaurant, error)
	AddRestaurant(ctx context.Context, restaurant restaurant.Restaurant, dishesToAdd []dish.Dish) error
	GetDishesForRestaurant(ctx context.Context, restaurantId string) ([]dish.Dish, error)
}

type RestaurantServiceImplementation struct {
	Logger         *zap.SugaredLogger
	RestaurantRepo database.RestaurantRepository
	DishRepo       database.DishRepository
}

func NewRestaurantService(logger *zap.SugaredLogger, restaurantRepo database.RestaurantRepository, dishRepo database.DishRepository) RestaurantService {
	return RestaurantServiceImplementation{
		Logger:         logger,
		RestaurantRepo: restaurantRepo,
		DishRepo:       dishRepo,
	}
}

func (s RestaurantServiceImplementation) GetAllRestaurants(ctx context.Context) ([]restaurant.Restaurant, error) {
	restaurants, err := s.RestaurantRepo.GetAllRestaurants(ctx)
	if err != nil {
		return []restaurant.Restaurant{}, err
	}
	return restaurants, nil
}

func (s RestaurantServiceImplementation) AddRestaurant(ctx context.Context, restaurantToAdd restaurant.Restaurant, dishesToAdd []dish.Dish) error {
	existingRestaurant, err := s.RestaurantRepo.GetRestaurantByName(ctx, restaurantToAdd.Name)

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

		err = s.RestaurantRepo.UpdateRestaurantLocations(ctx, *existingRestaurant, locationsToAdd)
		if err != nil {
			return err
		}
		return nil
	}

	err = s.DishRepo.AddDishes(ctx, dishesToAdd)
	if err != nil {
		return err
	}
	err = s.RestaurantRepo.AddRestaurant(ctx, restaurantToAdd)
	if err != nil {
		return err
	}

	return nil
}

func (s RestaurantServiceImplementation) GetDishesForRestaurant(ctx context.Context, restaurantId string) ([]dish.Dish, error) {
	restaurant, err := s.RestaurantRepo.GetRestaurantById(ctx, restaurantId)
	if err != nil {
		if errors.Is(err, database.ErrRestaurantNotFound) {
			return []dish.Dish{}, ErrRestaurantNotFound
		}
		return []dish.Dish{}, err
	}
	dishes := make([]dish.Dish, len(restaurant.DishIds))
	for i, dishId := range restaurant.DishIds {
		dishToRetrieve, err := s.DishRepo.GetDishById(ctx, dishId)
		if err != nil {
			return []dish.Dish{}, err
		}
		dishes[i] = *dishToRetrieve
	}
	return dishes, nil
}

func (s RestaurantServiceImplementation) locationsMatch(location restaurant.Location, targetLocation restaurant.Location) bool {
	return location.StreetAddressLine1 == targetLocation.StreetAddressLine1 && location.StreetAddressLine2 == targetLocation.StreetAddressLine2 && location.City == targetLocation.City && location.Province == targetLocation.Province && location.Country == targetLocation.Country && location.PostalCode == targetLocation.PostalCode
}
