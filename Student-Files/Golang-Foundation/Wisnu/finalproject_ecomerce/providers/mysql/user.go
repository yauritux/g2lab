package mysql

import (
	"finalproject_ecomerce/domain"
	"finalproject_ecomerce/engine"
	"time"

	"upper.io/db.v2/lib/sqlbuilder"
)

type (
	userRepository struct {
		sess sqlbuilder.Database
	}
)

func NewUserRepository(sess sqlbuilder.Database) engine.UserRepository {
	return &userRepository{sess: sess}
}

func (ur *userRepository) Add(u *domain.User) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	res, err := ur.sess.InsertInto(users).Values(u).Exec()
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	u.ID = uint(id)
	return err
}

func (ur *userRepository) One(id uint) (*domain.User, error) {
	var u domain.User
	return &u, handleErr(ur.sess.SelectFrom(users).Where(`id=?`, id).One(&u))
}

func (ur *userRepository) OneByEmail(email string) (*domain.User, error) {
	var u domain.User
	return &u, handleErr(ur.sess.SelectFrom(users).Where(`email=?`, email).One(&u))
}

func (ur *userRepository) ExistsByEmail(email string) (bool, error) {
	row, err := ur.sess.QueryRow(`SELECT COUNT(id) FROM users WHERE email=?`, email)
	if err != nil {
		return false, err
	}
	var n int
	err = row.Scan(&n)
	return n > 0, err
}

func (ur *userRepository) Update(u *domain.User) error {
	u.UpdatedAt = time.Now()
	_, err := ur.sess.Update(users).Set(u).Where(`id=?`, u.ID).Exec()
	return err
}
