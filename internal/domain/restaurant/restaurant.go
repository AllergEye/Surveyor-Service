package restaurant

import "github.com/google/uuid"

type Restaurant struct {
	RestaurantId uuid.UUID
	DishIds      []uuid.UUID
	Name         string
	Locations    []Location
}

func NewRestaurant(dishIds []uuid.UUID, name string, locations []Location) Restaurant {
	return Restaurant{
		RestaurantId: uuid.New(),
		DishIds:      dishIds,
		Name:         name,
		Locations:    locations,
	}
}
