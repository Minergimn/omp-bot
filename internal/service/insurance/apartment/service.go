package apartment

type ApartmentService interface {
	Describe(apartmentId uint64) (*Apartment, error)
	List(cursor uint64, limit uint64) ([]Apartment, error)
	Create(Apartment) (uint64, error)
	Update(apartmentId uint64, apartment Apartment) error
	Remove(apartmentId uint64) (bool, error)
}

type InsuranceApartmentService struct {}

func NewInsuranceApartmentService() *InsuranceApartmentService {
	return &InsuranceApartmentService{}
}

func (d *InsuranceApartmentService) Describe(apartmentId uint64) (*Apartment, error) {
	panic("implement me")
}

func (d *InsuranceApartmentService) List(cursor uint64, limit uint64) ([]Apartment, error) {
	panic("implement me")
}

func (d *InsuranceApartmentService) Create(apartment Apartment) (uint64, error) {
	panic("implement me")
}

func (d *InsuranceApartmentService) Update(apartmentId uint64, apartment Apartment) error {
	panic("implement me")
}

func (d *InsuranceApartmentService) Remove(apartmentId uint64) (bool, error) {
	panic("implement me")
}
