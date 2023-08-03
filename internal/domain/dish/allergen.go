package dish

import "strings"

const (
	PEANUTS   = "PEANUT"
	TREE_NUTS = "TREE_NUTS"
	SESAME    = "SESAME"
	MILK      = "MILK"
	EGGS      = "EGGS"
)

type Allergen struct {
	Name        string
	Probability int
}

func NewAllergen(name string) Allergen {
	return Allergen{
		Name:        strings.ToUpper(name),
		Probability: 100,
	}
}

func IsValidAllergen(name string) bool {
	str := strings.ToUpper(name)
	switch str {
	case PEANUTS, TREE_NUTS, SESAME, MILK, EGGS:
		return true
	default:
		return false
	}
}
