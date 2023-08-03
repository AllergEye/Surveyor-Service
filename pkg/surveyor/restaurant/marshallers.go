package surveyor_restaurant

import (
	"github.com/allergeye/surveyor-service/internal/domain/dish"
	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	"github.com/google/uuid"
)

type Marshallers interface {
	MarshalRestaurantRequestBody(restaurantRequestBody AddRestaurantRequestBody) (*restaurant.Restaurant, []dish.Dish, error)
	MarshalRestaurantLocationRequestBody(locationsRequestBody []AddRestaurantLocationRequestBody) ([]restaurant.Location, error)
	MarshalRestaurantDishRequestBody(dishesRequestBody []AddRestaurantDishRequestBody) ([]dish.Dish, error)
	MarshalRestaurantDishAllergenRequestBody(allergensRequestBody []AddRestaurantDishAllergenRequestBody) ([]dish.Allergen, error)
}

type MarshallersImplementation struct{}

func NewMarshallers() Marshallers {
	return MarshallersImplementation{}
}

func (m MarshallersImplementation) MarshalRestaurantRequestBody(restaurantRequestBody AddRestaurantRequestBody) (*restaurant.Restaurant, []dish.Dish, error) {
	locations, err := m.MarshalRestaurantLocationRequestBody(restaurantRequestBody.Locations)
	if err != nil {
		return nil, []dish.Dish{}, err
	}
	dishes, err := m.MarshalRestaurantDishRequestBody(restaurantRequestBody.Dishes)
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

func (m MarshallersImplementation) MarshalRestaurantLocationRequestBody(locationsRequestBody []AddRestaurantLocationRequestBody) ([]restaurant.Location, error) {
	locations := make([]restaurant.Location, len(locationsRequestBody))
	for i, locationRequest := range locationsRequestBody {
		location := restaurant.NewLocation(locationRequest.StreetAddressLine1, locationRequest.StreetAddressLine2, locationRequest.City, locationRequest.Province, locationRequest.Country, locationRequest.PostalCode)
		locations[i] = location
	}

	return locations, nil
}

func (m MarshallersImplementation) MarshalRestaurantDishRequestBody(dishesRequestBody []AddRestaurantDishRequestBody) ([]dish.Dish, error) {
	dishes := make([]dish.Dish, len(dishesRequestBody))
	for i, dishRequest := range dishesRequestBody {
		allergens, err := m.MarshalRestaurantDishAllergenRequestBody(dishRequest.Allergens)
		if err != nil {
			return []dish.Dish{}, err
		}
		dish := dish.NewDish(dishRequest.Name, allergens)
		dishes[i] = dish
	}

	return dishes, nil
}

func (m MarshallersImplementation) MarshalRestaurantDishAllergenRequestBody(allergensRequestBody []AddRestaurantDishAllergenRequestBody) ([]dish.Allergen, error) {
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
