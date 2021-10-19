package apartment

import (
	"fmt"
)

var allEntities = map[uint64]Apartment{
	1 : {ApartmentId: 1, Object: "Apartment one", Owner: "Owner One"},
	2 : {ApartmentId: 2, Object: "Apartment two", Owner: "Owner Two"},
	3 : {ApartmentId: 3, Object: "Apartment three", Owner: "Owner Three"},
	4 : {ApartmentId: 4, Object: "Apartment four", Owner: "Owner Four"},
	5 : {ApartmentId: 5, Object: "Apartment five", Owner: "Owner Five"},
	6 : {ApartmentId: 6, Object: "Apartment six", Owner: "Owner Six"},
	7 : {ApartmentId: 7, Object: "Apartment seven", Owner: "Owner Seven"},
	8 : {ApartmentId: 8, Object: "Apartment eight", Owner: "Owner Eight"},
	9 : {ApartmentId: 9, Object: "Apartment nine", Owner: "Owner Nine"},
	10 : {ApartmentId: 10, Object: "Apartment ten", Owner: "Owner Ten"},
}

type Apartment struct {
	ApartmentId uint64
	Object      string
	Owner       string
}

func (a *Apartment) String() string {
	return fmt.Sprintf("Apartment #%d title: %s owner: %s", a.ApartmentId, a.Object, a.Owner)
}
