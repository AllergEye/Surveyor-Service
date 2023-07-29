package restaurant

type Location struct {
	StreetAddressLine1 string
	StreetAddressLine2 string
	City               string
	Province           string
	Country            string
	PostalCode         string
}

func NewLocation(streetAddressLine1 string, streetAddressLine2 string, city string, province string, country string, postalCode string) *Location {
	return &Location{
		StreetAddressLine1: streetAddressLine1,
		StreetAddressLine2: streetAddressLine2,
		City:               city,
		Province:           province,
		Country:            country,
		PostalCode:         postalCode,
	}
}
