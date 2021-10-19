package apartment

import (
	"log"
	"sort"
)

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

	keys := make([]uint64, 0)
	for k, _ := range allEntities {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	if limit != 0 {
		var count uint64
		for i, k := range keys {
			if i < int(cursor){
				continue
			}
			result = append(result, allEntities[k])
			count++

			if count == limit{
				break
			}
		}
	} else {
		for _, k := range keys {
			result = append(result, allEntities[k])
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
