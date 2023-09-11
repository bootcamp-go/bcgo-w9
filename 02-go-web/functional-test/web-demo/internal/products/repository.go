package products

import "errors"

var (
	ErrProductNotFound = errors.New("product not found")
)

type Repository interface {
	GetByID(id int) (*Product, error)
}

type SliceBasedRepository struct {
	Data []*Product
}

func (repository SliceBasedRepository) GetByID(id int) (product *Product, err error) {
	for index := range repository.Data {
		if repository.Data[index].ID == id {
			product = repository.Data[index]
			return
		}
	}

	err = ErrProductNotFound
	return
}
