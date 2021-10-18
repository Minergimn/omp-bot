package apartment

import "fmt"

var allEntities = map[uint64]Apartment{
	0 : {ApartmentId: 0, Object: "Apartment one", Owner: "Owner One"},
	1 : {ApartmentId: 1, Object: "Apartment two", Owner: "Owner Two"},
	2 : {ApartmentId: 2, Object: "Apartment three", Owner: "Owner Three"},
	3 : {ApartmentId: 3, Object: "Apartment four", Owner: "Owner Four"},
	4 : {ApartmentId: 4, Object: "Apartment five", Owner: "Owner Five"},
}

type Apartment struct {
	ApartmentId uint64
	Object      string
	Owner       string
}

func (a *Apartment) String() string {
	return fmt.Sprintf("Apartment #%d title: %s owner: %s", a.ApartmentId, a.Object, a.Owner)
}
