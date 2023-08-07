package surveyor_dish

type AddDishesToRestaurantRequestBody struct {
	RestaurantId string
	Dishes       []AddDishesToRestaurantDishesRequestBody
}

type AddDishesToRestaurantDishesRequestBody struct {
	Name      string
	Allergens []AddDishesToRestaurantDishesAllergensRequestBody
}

type AddDishesToRestaurantDishesAllergensRequestBody struct {
	Name               string
	IsProbabilityKnown bool
	Probability        int
}
