package surveyor

import (
	"github.com/allergeye/surveyor-service/internal/domain/dish"
	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	"github.com/google/uuid"
)

func marshalRestaurantRequestBody(restaurantRequestBody AddRestaurantRequestBody) (*restaurant.Restaurant, []dish.Dish, error) {
	locations, err := marshalRestaurantLocationRequestBody(restaurantRequestBody.Locations)
	if err != nil {
		return nil, []dish.Dish{}, err
	}
	dishes, err := marshalRestaurantDishRequestBody(restaurantRequestBody.Dishes)
	if err != nil {
		return nil, []dish.Dish{}, err
	}
	dishIds := make([]uuid.UUID, len(dishes))
	for i, dish := range dishes {
		dishId := dish.DishId
		dishIds[i] = dishId
	}

	restaurant := restaurant.NewRestaurant(dishIds, restaurantRequestBody.Name, locations)
	return &restaurant, dishes, nil
}

func marshalRestaurantLocationRequestBody(locationsRequestBody []AddRestaurantLocationRequestBody) ([]restaurant.Location, error) {
	locations := make([]restaurant.Location, len(locationsRequestBody))
	for i, locationRequest := range locationsRequestBody {
		location := restaurant.NewLocation(locationRequest.StreetAddressLine1, locationRequest.StreetAddressLine2, locationRequest.City, locationRequest.Province, locationRequest.Country, locationRequest.PostalCode)
		locations[i] = location
	}

	return locations, nil
}

func marshalRestaurantDishRequestBody(dishesRequestBody []AddRestaurantDishRequestBody) ([]dish.Dish, error) {
	dishes := make([]dish.Dish, len(dishesRequestBody))
	for i, dishRequest := range dishesRequestBody {
		allergens, err := marshalRestaurantDishAllergenRequestBody(dishRequest.Allergens)
		if err != nil {
			return []dish.Dish{}, err
		}
		dish := dish.NewDish(dishRequest.Name, allergens)
		dishes[i] = dish
	}

	return dishes, nil
}

func marshalRestaurantDishAllergenRequestBody(allergensRequestBody []AddRestaurantDishAllergenRequestBody) ([]dish.Allergen, error) {
	allergens := make([]dish.Allergen, len(allergensRequestBody))
	for i, allergenRequest := range allergensRequestBody {
		if !dish.IsValidAllergen(allergenRequest.Name) {
			return []dish.Allergen{}, ErrInvalidAllergen
		}
		allergen := dish.NewAllergen(allergenRequest.Name)
		allergens[i] = allergen
	}

	return allergens, nil
}
