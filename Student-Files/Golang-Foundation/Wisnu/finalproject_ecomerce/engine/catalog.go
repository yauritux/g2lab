package engine

import (
	"finalproject_ecomerce/domain"
	"sync"
)

type (
	Catalog interface {
		AddProduct(*AddProductRequest) (*domain.Product, error)
		ShowProduct(*ShowProductRequest) (*domain.Product, error)
		ListProducts(*ListProductsRequest) ([]*domain.Product, error)
		UpdateProduct(*UpdateProductRequest) error
		DeleteProduct(*DeleteProductRequest) error
	}

	catalog struct {
		validator Validator
		repo      CatalogRepository
		imgRepo   ImageRepository
	}
)

var (
	catalogInstance Catalog
	catalogOnce     sync.Once
)

func (f *factory) NewCatalog() Catalog {
	catalogOnce.Do(func() {
		catalogInstance = &catalog{
			repo:      f.NewCatalogRepository(),
			imgRepo:   f.NewImageRepository(),
			validator: f.v,
		}
	})
	return catalogInstance
}
