package restaurant

import "github.com/google/uuid"

type Restaurant struct {
	RestaurantId uuid.UUID
	Name         string
	Locations    []Location
}

func NewRestaurant(restaurantId uuid.UUID, name string, locations []Location) *Restaurant {
	return &Restaurant{
		RestaurantId: restaurantId,
		Name:         name,
		Locations:    locations,
	}
}
