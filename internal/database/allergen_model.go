package database

type ALLERGEN_ENUM string

const (
	PEANUTS   ALLERGEN_ENUM = "PEANUT"
	TREE_NUTS ALLERGEN_ENUM = "TREE_NUTS"
	SESAME    ALLERGEN_ENUM = "SESAME"
	MILK      ALLERGEN_ENUM = "MILK"
	EGGS      ALLERGEN_ENUM = "EGGS"
)

type AllergenModel struct {
	Name        ALLERGEN_ENUM
	Probability int
}
