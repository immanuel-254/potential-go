package models

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64     `db:"id" json:"id"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"password"`
	Active   bool      `db:"active" json:"active"`
	Staff    bool      `db:"staff" json:"staff"`
	Admin    bool      `db:"admin" json:"admin"`
	Created  time.Time `db:"created" json:"created"`
	Updated  time.Time `db:"updated" json:"updated"`
}

func (u *User) UserCreate(db *sqlx.DB, w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return errors.New("method not allowed")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	var user User
	query := "INSERT INTO users (email, password, active, staff, admin, created) VALUES (?, ?, ?, ?, ?, ?) RETURNING id, email, created, updated;"

	row := db.QueryRow(query,
		u.Email,
		string(hash),
		u.Active,
		u.Staff,
		u.Admin,
		u.Created,
	)

	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.Created,
		&user.Updated,
	)

	if err != nil {
		return err
	}

	var data struct {
		User `json:"user"`
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	return nil
}

func (u *User) UserRead(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {

}

func (u *User) UserUpdate(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {

}

func (u *User) UserDelete(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {

}
