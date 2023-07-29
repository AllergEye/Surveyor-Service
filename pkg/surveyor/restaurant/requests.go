package surveyor

import "github.com/allergeye/surveyor-service/internal/domain/restaurant"

type AddRestaurantRequestBody struct {
	Name      string
	Locations []restaurant.Location
}
