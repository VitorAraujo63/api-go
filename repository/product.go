package repository

import(
	"errors"
	"minha-api-go/models"
	"sync"
)

type ProductRepository struct {
	mu sync.RWMutex
	products map[string]models.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: make(map[string]models.Product),
	}
}

func (r *ProductRepository) Save(p models.Product) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.products[p.ID] = p
}

func (r *ProductRepository) FindAll() []models.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]models.Product, 0, len(r.products))
	for _, p := range r.products {
		list = append(list, p)
	}
	return list
}

func (r *ProductRepository) FindByID(id string) (models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, exists := r.products[id]
	if !exists {
		return models.Product{}, errors.New("Produto não encontrado.")
	}
	return p, nil
}

func (r *ProductRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[id]; !exists {
		return errors.New("Produto não encontrado.")
	}
	delete(r.products, id)
	return nil
}