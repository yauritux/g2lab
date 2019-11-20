package web

import (
	"errors"
	"finalproject_ecomerce/engine"
	"net/http"
	"strings"

	"strconv"

	"gopkg.in/alioygur/gores.v1"
)

type (
	catalog struct {
		engine.Catalog
	}
)

func newCatalog(f engine.Factory) *catalog {
	return &catalog{f.NewCatalog()}
}

func (c *catalog) addProduct(w http.ResponseWriter, r *http.Request) error {
	ar := new(engine.AddProductRequest)
	if err := decodeReq(r, ar); err != nil {
		return err
	}
	p, err := c.AddProduct(ar)
	if err != nil {
		return err
	}
	return gores.JSON(w, http.StatusCreated, response{p})
}

func (c *catalog) showProduct(w http.ResponseWriter, r *http.Request) error {
	id := urlParamMustInt("id", r)
	p, err := c.ShowProduct(&engine.ShowProductRequest{ID: uint(id)})
	if err != nil {
		return err
	}
	return gores.JSON(w, http.StatusOK, response{p})
}

func (c *catalog) listProducts(w http.ResponseWriter, r *http.Request) error {
	req := new(engine.ListProductsRequest)
	req.InCategories = categoryQueryValue(r)

	req.Sort = engine.SortByIDDesc

	limit, err := queryValueInt("limit", r)
	if err != nil {
		return newWebErr(badRequestErrCode, http.StatusBadRequest, errors.New("limit value must be integer"))
	}
	offset, err := queryValueInt("offset", r)
	if err != nil {
		return newWebErr(badRequestErrCode, http.StatusBadRequest, errors.New("offset value must be integer"))
	}

	req.Limit = limit
	req.Offset = offset

	ps, err := c.ListProducts(req)
	if err != nil {
		return err
	}
	return gores.JSON(w, http.StatusOK, response{ps})
}

func (c *catalog) updateProduct(w http.ResponseWriter, r *http.Request) error {
	req := new(engine.UpdateProductRequest)
	if err := decodeReq(r, req); err != nil {
		return err
	}

	id := urlParamMustInt("id", r)
	req.ID = uint(id)
	if err := c.UpdateProduct(req); err != nil {
		return err
	}

	gores.NoContent(w)
	return nil
}

func (c *catalog) deleteProduct(w http.ResponseWriter, r *http.Request) error {
	req := new(engine.DeleteProductRequest)
	if err := decodeReq(r, req); err != nil {
		return err
	}

	id := urlParamMustInt("id", r)
	req.ID = uint(id)
	if err := c.DeleteProduct(req); err != nil {
		return err
	}

	gores.NoContent(w)
	return nil
}

// categoryQueryValue gets category query values like 1,2,3 as []uint{1, 2, 3}
func categoryQueryValue(r *http.Request) []uint {
	var cs []uint
	c := queryValue("category", r)
	if c != "" {
		for _, id := range strings.Split(c, ",") {
			if id == "" {
				continue
			}

			intID, err := strconv.Atoi(strings.Trim(id, " "))
			if err != nil {
				continue
			}
			cs = append(cs, uint(intID))
		}
	}
	return cs
}
