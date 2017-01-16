package models

import (
	"errors"

	"github.com/patricklecuyer/cr460-lab1/config"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Contact is a type representing a contact
type Contact struct {
	_id         string `bson:"_id"`
	FirstName   string `json:"firstname" bson:"firstname" binding:"required"`
	LastName    string `json:"lastname" bson:"lastname" binding:"required"`
	Email       string `json:"email" bson:"email" binding:"required" `
	PhoneNumber string `json:"phonenumber" bson:"phonenumber" validate:"basic"`
}

func compare(a, b Contact) bool {
	if a.FirstName != b.FirstName {
		return false
	}

	if a.LastName != b.LastName {
		return false
	}

	if a.Email != b.Email {
		return false
	}
	if a.PhoneNumber != b.PhoneNumber {
		return false
	}

	return true
}

// ErrContactAlreadyExists user already exists
var ErrContactAlreadyExists = errors.New("cr460: Contact already exists")

// ContactByEmail Get returns a user based on email
func ContactByEmail(email string) (*Contact, error) {

	db := config.AppConfig.Datastore.GetDBSession()
	defer db.Session.Close()
	var c Contact
	err := db.C("Contacts").Find(bson.M{"email": email}).One(&c)
	return &c, err

}

// GetAllContacts gets all contacts
func GetAllContacts() ([]Contact, error) {

	db := config.AppConfig.Datastore.GetDBSession()
	defer db.Session.Clone()
	c := make([]Contact, 1)
	err := db.C("Contacts").Find(bson.M{}).All(&c)
	return c, err
}

// Update a user
func (c *Contact) Update(n Contact) error {

	db := config.AppConfig.Datastore.GetDBSession()
	db.Session.SetSafe(&mgo.Safe{})
	defer db.Session.Close()

	err := db.C("Contacts").Update(bson.M{"email": c.Email}, n)
	return err
}

// Insert a new user
func (c *Contact) Insert() error {

	if _, err := ContactByEmail(c.Email); err == nil {
		return ErrContactAlreadyExists
	}

	db := config.AppConfig.Datastore.GetDBSession()
	defer db.Session.Close()

	err := db.C("Contacts").Insert(c)
	return err

}
