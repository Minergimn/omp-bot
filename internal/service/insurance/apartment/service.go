package apartment

import "log"

type ApartmentService interface {
	Describe(apartmentId uint64) (*Apartment, error)
	List(cursor uint64, limit uint64) ([]Apartment, error)
	Create(Apartment) (uint64, error)
	Update(apartmentId uint64, apartment Apartment) error
	Remove(apartmentId uint64) (bool, error)
}

type InsuranceApartmentService struct{}

func NewInsuranceApartmentService() *InsuranceApartmentService {
	return &InsuranceApartmentService{}
}

func (d *InsuranceApartmentService) Describe(apartmentId uint64) (*Apartment, error) {
	if value, ok := allEntities[apartmentId]; ok {
		return &value, nil
	} else {
		return nil, nil
	}
}

func (d *InsuranceApartmentService) List(cursor uint64, limit uint64) ([]Apartment, error) {
	result := make([]Apartment, 0)

	if limit != 0 {
		stopIndex := cursor + limit

		for i := cursor; i < stopIndex; i++ {
			if value, ok := allEntities[i]; ok {
				result = append(result, value)
			} else {
				break
			}
		}
	} else {
		for _, a := range allEntities {
			result = append(result, a)
		}
	}

	return result, nil

}

func (d *InsuranceApartmentService) Create(apartment Apartment) (uint64, error) {
	panic("implement me")
}

func (d *InsuranceApartmentService) Update(apartmentId uint64, apartment Apartment) error {
	panic("implement me")
}

func (d *InsuranceApartmentService) Remove(apartmentId uint64) (bool, error) {
	if _, ok := allEntities[apartmentId]; ok{
		delete(allEntities, apartmentId)
		return ok, nil
	} else {
		log.Printf("Apartment %d doesn't exist", apartmentId)
		return ok, nil
	}
}
