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
	Name               string
	IsProbabilityKnown bool
	Probability        int
}

func NewAllergen(name string, isProbabilityKnown bool, probability int) Allergen {
	return Allergen{
		Name:               strings.ToUpper(name),
		IsProbabilityKnown: isProbabilityKnown,
		Probability:        probability,
	}
}

func IsValidAllergenName(name string) bool {
	str := strings.ToUpper(name)
	switch str {
	case PEANUTS, TREE_NUTS, SESAME, MILK, EGGS:
		return true
	default:
		return false
	}
}
