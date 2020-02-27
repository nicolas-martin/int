package repository

import (
	"github.com/nicolas-martin/int/internal/types"
)

// IRepository represents the interface of the repository
type IRepository interface {
	Create(d *types.Deposit) error
	RetrieveAll(id int) ([]*types.Deposit, error)
}

// Repository implements IRepository
type Repository struct {
	data []*types.Deposit
}

var _ IRepository = &Repository{}

// NewRepository returns an IRepository
func NewRepository() IRepository {
	d := make([]*types.Deposit, 0)
	return &Repository{data: d}
}

// Create creates a deposit entry in the store
func (r *Repository) Create(d *types.Deposit) error {
	r.data = append(r.data, d)
	return nil
}

// RetrieveAll retrieves all the deposis for a user
func (r *Repository) RetrieveAll(id int) ([]*types.Deposit, error) {
	result := make([]*types.Deposit, 0)
	for _, v := range r.data {
		if v.CustomerID == id {
			result = append(result, v)
		}
	}

	return result, nil

}
