package database

import (
	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func unmarshalRestaurant(rm *RestaurantModel) (*restaurant.Restaurant, error) {
	restaurantId, err := uuid.Parse(rm.RestaurantId)
	if err != nil {
		return nil, err
	}

	locations, err := unmarshalLocations(rm.Locations)
	if err != nil {
		return nil, err
	}

	return restaurant.NewRestaurant(restaurantId, rm.Name, locations), nil
}

func unmarshalLocations(locations []LocationModel) ([]restaurant.Location, error) {
	unmarshalledLocations := make([]restaurant.Location, len(locations))
	for i, location := range locations {
		l := restaurant.NewLocation(location.StreetAddressLine1, location.StreetAddressLine2, location.City, location.Province, location.Country, location.PostalCode)
		unmarshalledLocations[i] = *l
	}

	return unmarshalledLocations, nil
}

func marshalRestaurant(restaurant restaurant.Restaurant) (*RestaurantModel, error) {
	locations, err := marshalLocations(restaurant.Locations)
	if err != nil {
		return nil, err
	}

	return &RestaurantModel{
		ID:           primitive.NewObjectID(),
		RestaurantId: restaurant.RestaurantId.String(),
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
