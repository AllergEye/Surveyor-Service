package database

import (
	"github.com/allergeye/surveyor-service/internal/domain/dish"
	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func unmarshalRestaurant(rm *RestaurantModel) (*restaurant.Restaurant, error) {
	locations, err := unmarshalLocations(rm.Locations)
	if err != nil {
		return nil, err
	}

	return &restaurant.Restaurant{
		RestaurantId: rm.RestaurantId,
		DishIds:      rm.DishIds,
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

func unmarshalDish(dm *DishModel) (*dish.Dish, error) {
	allergens, err := unmarshalAllergens(dm.Allergens)
	if err != nil {
		return nil, err
	}

	return &dish.Dish{
		DishId:    dm.DishId,
		Name:      dm.Name,
		Allergens: allergens,
	}, nil
}

func unmarshalAllergens(allergens []AllergenModel) ([]dish.Allergen, error) {
	unmarshalledAllergens := make([]dish.Allergen, len(allergens))
	for i, allergen := range allergens {
		a := dish.Allergen{
			Name:        string(allergen.Name),
			Probability: allergen.Probability,
		}
		unmarshalledAllergens[i] = a
	}

	return unmarshalledAllergens, nil
}

func marshalRestaurant(restaurant restaurant.Restaurant) (*RestaurantModel, error) {
	locations, err := marshalLocations(restaurant.Locations)
	if err != nil {
		return nil, err
	}

	return &RestaurantModel{
		ID:           primitive.NewObjectID(),
		RestaurantId: restaurant.RestaurantId,
		DishIds:      restaurant.DishIds,
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
		DishId:    dish.DishId,
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
