package database

import (
	"github.com/allergeye/surveyor-service/internal/domain/dish"
	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func unmarshalRestaurant(rm *RestaurantModel) (*restaurant.Restaurant, error) {
	restaurantId, err := uuid.Parse(rm.RestaurantId)
	if err != nil {
		return nil, err
	}
	dishIds := make([]uuid.UUID, len(rm.DishIds))
	for i, dishId := range rm.DishIds {
		uuidDishId, err := uuid.Parse(dishId)
		if err != nil {
			return nil, err
		}
		dishIds[i] = uuidDishId
	}

	locations, err := unmarshalLocations(rm.Locations)
	if err != nil {
		return nil, err
	}

	return &restaurant.Restaurant{
		RestaurantId: restaurantId,
		DishIds:      dishIds,
		Name:         rm.Name,
		Locations:    locations,
	}, nil
}

func unmarshalLocations(locations []LocationModel) ([]restaurant.Location, error) {
	unmarshalledLocations := make([]restaurant.Location, len(locations))
	for i, location := range locations {
		l := restaurant.Location{
			StreetAddressLine1: location.StreetAddressLine1,
			StreetAddressLine2: location.StreetAddressLine2,
			City:               location.City,
			Province:           location.Province,
			Country:            location.Country,
			PostalCode:         location.PostalCode,
		}
		unmarshalledLocations[i] = l
	}

	return unmarshalledLocations, nil
}

func marshalRestaurant(restaurant restaurant.Restaurant) (*RestaurantModel, error) {
	locations, err := marshalLocations(restaurant.Locations)
	if err != nil {
		return nil, err
	}

	dishIds := make([]string, len(restaurant.DishIds))
	for i, dishId := range restaurant.DishIds {
		dishIds[i] = dishId.String()
	}

	return &RestaurantModel{
		ID:           primitive.NewObjectID(),
		RestaurantId: restaurant.RestaurantId.String(),
		DishIds:      dishIds,
		Name:         restaurant.Name,
		Locations:    locations,
	}, nil
}

func marshalLocations(locations []restaurant.Location) ([]LocationModel, error) {
	marshalledLocations := make([]LocationModel, len(locations))

	for i, location := range locations {
		l := LocationModel{
			StreetAddressLine1: location.StreetAddressLine1,
			StreetAddressLine2: location.StreetAddressLine2,
			City:               location.City,
			Province:           location.Province,
			Country:            location.Country,
			PostalCode:         location.PostalCode,
		}
		marshalledLocations[i] = l
	}

	return marshalledLocations, nil
}

func marshalDish(dish dish.Dish) (*DishModel, error) {
	allergens, err := marshalAllergens(dish.Allergens)
	if err != nil {
		return nil, err
	}

	return &DishModel{
		ID:        primitive.NewObjectID(),
		DishId:    dish.DishId.String(),
		Name:      dish.Name,
		Allergens: allergens,
	}, nil
}

func marshalAllergens(allergens []dish.Allergen) ([]AllergenModel, error) {
	marshalledAllergens := make([]AllergenModel, len(allergens))

	for i, allergen := range allergens {
		a := AllergenModel{
			Name:        ALLERGEN_ENUM(allergen.Name),
			Probability: allergen.Probability,
		}

		marshalledAllergens[i] = a
	}

	return marshalledAllergens, nil
}
