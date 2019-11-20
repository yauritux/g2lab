package web

import (
	"finalproject_ecomerce/engine"
	"net/http"

	gores "gopkg.in/alioygur/gores.v1"

)

type (
	errHandlerFunc func(http.ResponseWriter, *http.Request) error

	errResponse struct {
		Code     errCode `json:"code"`
		HTTPCode int     `json:"httpCode"`
		Error    string  `json:"error"`
		Inner    string  `json:"inner,omitempty"`
	}
)

func (h errHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		h.handle(w, err)
	}
}

func (h errHandlerFunc) handle(w http.ResponseWriter, err error) {
	httpErr := func() *webErr {

		if err == engine.ErrNoRows {
			return newWebErr(noRowsErrCode, http.StatusNotFound, err)
		}

		switch t := err.(type) {
		case *webErr:
			return t
		case *engine.ValidationErr:
			return newWebErr(validationErrCode, http.StatusBadRequest, err)
		}
		// default
		return newWebErr(unknownErrCode, http.StatusInternalServerError, err)
	}()

	errRes := errResponse{Code: httpErr.Code, HTTPCode: httpErr.HTTPCode, Error: httpErr.Error()}
	if httpErr.inner != nil {
		errRes.Inner = httpErr.inner.Error()
	}
	gores.JSON(w, httpErr.HTTPCode, errRes)
}
