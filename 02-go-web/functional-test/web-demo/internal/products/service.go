package products

type Service interface {
	GetByID(id int) (*Product, error)
}

type DefaultService struct {
	Repository Repository
}

func (service DefaultService) GetByID(id int) (*Product, error) {
	return service.Repository.GetByID(id)
}
