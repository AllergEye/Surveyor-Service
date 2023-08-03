package surveyor_restaurant

type AddRestaurantRequestBody struct {
	Name      string
	Locations []AddRestaurantLocationRequestBody
	Dishes    []AddRestaurantDishRequestBody
}

type AddRestaurantLocationRequestBody struct {
	StreetAddressLine1 string
	StreetAddressLine2 string
	City               string
	Province           string
	Country            string
	PostalCode         string
}

type AddRestaurantDishRequestBody struct {
	Name      string
	Allergens []AddRestaurantDishAllergenRequestBody
}

type AddRestaurantDishAllergenRequestBody struct {
	Name        string
	Probability int
}
