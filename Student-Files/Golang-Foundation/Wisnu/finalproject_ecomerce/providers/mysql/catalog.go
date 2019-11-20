package mysql

import (
	"finalproject_ecomerce/domain"
	"finalproject_ecomerce/engine"
	"time"

	"upper.io/db.v2/lib/sqlbuilder"
)

type (
	catalogRepository struct {
		repository
	}
)

func NewCatalogRepository(sess sqlbuilder.Database) engine.CatalogRepository {
	var r catalogRepository
	r.sess = sess
	return &r
}

func (c *catalogRepository) AddProduct(p *domain.Product) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	r, err := c.sess.InsertInto(products).Values(*p).Exec()
	if err != nil {
		return err
	}
	id, err := r.LastInsertId()
	p.ID = uint(id)
	return err
}

func (c *catalogRepository) OneProduct(id uint) (*domain.Product, error) {
	var p domain.Product
	return &p, handleErr(c.sess.SelectFrom(products).One(&p))
}

func (c *catalogRepository) OneActiveProduct(id uint) (*domain.Product, error) {
	var p domain.Product
	return &p, handleErr(c.sess.SelectFrom(products).Where(`is_active=?`, true).One(&p))
}

func (c *catalogRepository) FindActiveProducts(s engine.Sorting, offset int, limit int) ([]*domain.Product, error) {
	var ps []*domain.Product
	var orderBy = `id DESC`
	return ps, c.sess.SelectFrom(products).Where("is_active=?", true).OrderBy(orderBy).Offset(offset).Limit(limit).All(&ps)
}

func (c *catalogRepository) FindActiveProductsInCategories(ids []uint, s engine.Sorting, offset int, limit int) ([]*domain.Product, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var ps []*domain.Product
	var orderBy = `p.id DESC`

	return ps, c.sess.Select("p.*").
		From(pivotProductCategory+" AS pivot").
		Join(products+" AS p").
		On("pivot.product_id = p.id").
		Where("pivot.category_id IN ?", ids).
		Where("p.is_active=?", true).
		GroupBy("pivot.product_id").
		OrderBy(orderBy).
		Offset(offset).
		Limit(limit).
		All(&ps)
}

func (c *catalogRepository) FindProductsInCategories(ids []uint, s engine.Sorting, offset int, limit int) ([]*domain.Product, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var ps []*domain.Product
	var orderBy = `p.id DESC`

	return ps, c.sess.Select("p.*").
		From(pivotProductCategory+" AS pivot").
		Join(products+" AS p").
		On("pivot.product_id = p.id").
		Where("pivot.category_id IN ?", ids).
		GroupBy("pivot.product_id").
		OrderBy(orderBy).
		Offset(offset).
		Limit(limit).
		All(&ps)
}

// todo: update associations
func (c *catalogRepository) UpdateProduct(p *domain.Product) error {
	p.UpdatedAt = time.Now()
	_, err := c.sess.Update(products).Set(p).Where(`id=?`, p.ID).Exec()
	return err
}

// DeleteWithAssoc deletes product with its associations
func (c *catalogRepository) DeleteProductWithAssoc(id uint) error {
	return c.sess.Tx(func(tx sqlbuilder.Tx) error {
		if _, err := tx.DeleteFrom(products).Where(`id=?`, id).Exec(); err != nil {
			return err
		}

		if _, err := tx.DeleteFrom(pivotProductCategory).Where(`product_id=?`, id).Exec(); err != nil {
			return err
		}
		return nil
	})
}

func (c *catalogRepository) FindCategoriesByIDs(ids []uint) ([]*domain.Category, error) {
	var cs []*domain.Category
	return cs, c.sess.SelectFrom(categories).Where(`id IN ?`, ids).All(&cs)
}
