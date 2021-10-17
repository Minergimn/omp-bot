package apartment

import "fmt"

var allEntities = []Apartment{
	{Title: "one", ApartmentId: 1},
	{Title: "two", ApartmentId: 2},
	{Title: "three", ApartmentId: 3},
	{Title: "four", ApartmentId: 4},
	{Title: "five", ApartmentId: 5},
}

type Apartment struct {
	ApartmentId uint64
	Title       string
}

func (a *Apartment) String() string{
	return fmt.Sprintf("Apartment #%d title: %s", a.ApartmentId, a.Title)
}
