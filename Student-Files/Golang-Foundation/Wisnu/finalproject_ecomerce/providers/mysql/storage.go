package mysql

import (
	"database/sql"
	"finalproject_ecomerce/engine"

	db "upper.io/db.v2"
	"upper.io/db.v2/lib/sqlbuilder"
)

type (
	tableName string

	repository struct {
		sess sqlbuilder.Database
	}
)

const (
	users                = `users`
	products             = `products`
	categories           = `categories`
	pivotProductCategory = `pivot_product_category`
)

func handleErr(err error) error {
	if err == db.ErrNoMoreRows || err == sql.ErrNoRows {
		return engine.ErrNoRows
	}
	return err
}
