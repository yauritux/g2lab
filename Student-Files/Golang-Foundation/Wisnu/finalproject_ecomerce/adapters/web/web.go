package web

import (
	"finalproject_ecomerce/engine"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func NewWebAdapter(f engine.Factory) http.Handler {
	r := mux.NewRouter()

	// base handler
	base := alice.New(newSetUserMid(f.NewUser()))
	// handler with auth required
	authRequired := base.Append(newAuthRequiredMid)
	// handler with admin only
	adminOnly := authRequired.Append(newAdminOnlyMid)

	user := newUser(f)
	catalog := newCatalog(f)

	r.Handle("/v1/auth/login", base.Then(errHandlerFunc(user.login))).Methods("POST")
	r.Handle("/v1/auth/register", base.Then(errHandlerFunc(user.register))).Methods("POST")
	r.Handle("/v1/auth/activate", base.Then(errHandlerFunc(user.activate))).Methods("POST")

	r.Handle("/v1/password/forgot", base.Then(errHandlerFunc(user.forgotPassword))).Methods("POST")
	r.Handle("/v1/password/reset", base.Then(errHandlerFunc(user.resetPassword))).Methods("POST")

	r.Handle("/v1/me", authRequired.Then(errHandlerFunc(user.me))).Methods("GET")
	r.Handle("/v1/me", authRequired.Then(errHandlerFunc(user.updateMe))).Methods("PATCH")

	r.Handle("/v1/admin/products", adminOnly.Then(errHandlerFunc(catalog.addProduct))).Methods("POST")
	r.Handle("/v1/admin/products/{id}", adminOnly.Then(errHandlerFunc(catalog.updateProduct))).Methods("PATCH")
	r.Handle("/v1/admin/products/{id}", adminOnly.Then(errHandlerFunc(catalog.deleteProduct))).Methods("DELETE")

	r.Handle("/v1/products", base.Then(errHandlerFunc(catalog.listProducts))).Methods("GET")
	r.Handle("/v1/products/{id}", base.Then(errHandlerFunc(catalog.showProduct))).Methods("GET")
	return r
}
