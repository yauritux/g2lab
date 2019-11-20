package engine

import (
	"finalproject_ecomerce/domain"
)

type (
	Sorting uint

	UserRepository interface {
		Add(*domain.User) error
		One(uint) (*domain.User, error)
		OneByEmail(string) (*domain.User, error)
		ExistsByEmail(string) (bool, error)
		Update(*domain.User) error
	}

	CatalogRepository interface {
		AddProduct(*domain.Product) error
		OneProduct(id uint) (*domain.Product, error)
		OneActiveProduct(id uint) (*domain.Product, error)
		FindActiveProducts(s Sorting, offset int, limit int) ([]*domain.Product, error)
		FindActiveProductsInCategories(ids []uint, s Sorting, offset int, limit int) ([]*domain.Product, error)
		FindProductsInCategories(ids []uint, s Sorting, offset int, limit int) ([]*domain.Product, error)
		UpdateProduct(*domain.Product) error
		DeleteProductWithAssoc(id uint) error
		FindCategoriesByIDs(ids []uint) ([]*domain.Category, error)
	}

	ImageRepository interface {
		FirstOrInit(string) (*domain.Image, error)
	}

	StorageFactory interface {
		NewUserRepository() UserRepository
		NewCatalogRepository() CatalogRepository
		NewImageRepository() ImageRepository
	}
)

const (
	SortByIDDesc Sorting = iota
)
