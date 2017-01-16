package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/patricklecuyer/planifio-api/config"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User is a type representing a user in the system
type User struct {
	_id          string `bson:"_id"`
	FirstName    string `json:"firstname" bson:"firstname" binding:"required" validate:"nonzero,basic"`
	LastName     string `json:"lastname" bson:"lastname" binding:"required" validate:"nonzero,basic"`
	Email        string `json:"email" bson:"email" binding:"required" validate:"basic"`
	PasswordHash []byte `bson:"passwordhash"`
}

func compare(a, b User) bool {
	if a.FirstName != b.FirstName {
		return false
	}

	if a.LastName != b.LastName {
		return false
	}

	if a.Email != b.Email {
		return false
	}
	return true
}

// ErrUserAlreadyExists user already exists
var ErrUserAlreadyExists = errors.New("planifio-api: User already exists")

//ErrUserNoPassword User does not have a password!
var ErrUserNoPassword = errors.New("User does not have a password!")

// UserByEmail Get returns a user based on email
func UserByEmail(email string) (*User, error) {

	db := config.AppConfig.Datastore.GetDBSession()
	defer db.Session.Close()
	var u User
	err := db.C("Users").Find(bson.M{"email": email}).One(&u)
	return &u, err

}

// SetPassword hashes the user's password
func (u *User) SetPassword(pw string) (err error) {
	u.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return
}

// ValidatePassword validates the password
func (u *User) ValidatePassword(p string) (err error) {
	return bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(p))
}

// Update a user
func (u *User) Update() error {

	db := config.AppConfig.Datastore.GetDBSession()
	db.Session.SetSafe(&mgo.Safe{})
	defer db.Session.Close()

	err := db.C("Users").Update(bson.M{"email": u.Email}, u)
	return err
}

// Insert a new user
func (u *User) Insert() error {

	if _, err := UserByEmail(u.Email); err == nil {
		return ErrUserAlreadyExists
	}

	if u.PasswordHash == nil {
		return ErrUserNoPassword
	}

	db := config.AppConfig.Datastore.GetDBSession()
	defer db.Session.Close()

	err := db.C("Users").Insert(u)
	return err

}
