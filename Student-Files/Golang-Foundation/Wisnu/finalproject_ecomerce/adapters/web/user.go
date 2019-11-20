package web

import (
	"finalproject_ecomerce/domain"
	"finalproject_ecomerce/engine"
	"gopkg.in/alioygur/gores.v1"
	"net/http"
)

type (
	user struct {
		engine.User
	}
)

func newUser(f engine.Factory) *user {
	return &user{f.NewUser()}
}

func (u *user) activate(w http.ResponseWriter, r *http.Request) error {
	req := new(engine.UserActivateRequest)
	if err := decodeReq(r, req); err != nil {
		return err
	}

	if err := u.Activate(req); err != nil {
		tokenErr, ok := err.(*engine.TokenErr)
		if ok {
			if tokenErr.Expired() {
				return newWebErr(expiredTokenErrCode, http.StatusBadRequest, err)
			}
			return newWebErr(invalidTokenErrCode, http.StatusBadRequest, err)
		}
		return err
	}

	gores.NoContent(w)
	return nil
}

func (u *user) login(w http.ResponseWriter, r *http.Request) error {
	req := new(engine.LoginRequest)
	if err := decodeReq(r, req); err != nil {
		return err
	}

	usr, err := u.Login(req)
	if err != nil {
		switch err {
		case engine.ErrWrongCredentials:
			return newWebErr(wrongCredErrCode, http.StatusUnauthorized, err)
		case engine.ErrInActiveUser:
			return newWebErr(inactiveUserErrCode, http.StatusUnauthorized, err)
		}
		return err
	}

	jwt, err := u.GenToken(usr, engine.AuthToken)
	if err != nil {
		return err
	}

	return gores.JSON(w, http.StatusOK, response{jwt})
}

func (u *user) register(w http.ResponseWriter, r *http.Request) error {
	req := new(engine.RegisterRequest)
	if err := decodeReq(r, req); err != nil {
		return err
	}

	usr, err := u.Register(req)
	if err != nil {
		if err == engine.ErrEmailExists {
			return newWebErr(emailExistsErrCode, http.StatusConflict, err)
		}
		return err
	}

	jwt, err := u.GenToken(usr, engine.AuthToken)
	if err != nil {
		return err
	}
	return gores.JSON(w, http.StatusCreated, response{jwt})
}

func (u *user) forgotPassword(w http.ResponseWriter, r *http.Request) error {
	req := new(engine.ForgotPasswordRequest)
	if err := decodeReq(r, req); err != nil {
		return err
	}

	if err := u.SendPasswordResetMail(req); err != nil {
		if err == engine.ErrNoRows {
			return newWebErr(noRowsErrCode, http.StatusNotFound, err)
		}
		return err
	}

	gores.NoContent(w)
	return nil
}

func (u *user) resetPassword(w http.ResponseWriter, r *http.Request) error {
	req := new(engine.ResetPasswordRequest)
	if err := decodeReq(r, req); err != nil {
		return err
	}

	if err := u.ResetPassword(req); err != nil {
		_, ok := err.(*engine.TokenErr)
		if ok {
			return newWebErr(invalidTokenErrCode, http.StatusBadRequest, err)
		}
		return err
	}

	gores.NoContent(w)
	return nil
}

func (u *user) me(w http.ResponseWriter, r *http.Request) error {
	me := domain.UserMustFromContext(r.Context())
	req := engine.ShowUserRequest{ID: me.ID}
	usr, err := u.Show(&req)
	if err != nil {
		return err
	}
	return gores.JSON(w, http.StatusOK, response{usr})
}

func (u *user) updateMe(w http.ResponseWriter, r *http.Request) error {
	me := domain.UserMustFromContext(r.Context())
	req := new(engine.UpdateUserRequest)
	if err := decodeReq(r, req); err != nil {
		return err
	}

	req.ID = me.ID

	if err := u.Update(req); err != nil {
		if err == engine.ErrEmailExists {
			return newWebErr(emailExistsErrCode, http.StatusUnprocessableEntity, err)
		}
		return err
	}

	gores.NoContent(w)
	return nil
}
