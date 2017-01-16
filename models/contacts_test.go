package models

import (
	"testing"

	"github.com/patricklecuyer/cr460-lab1/config"
	"github.com/stretchr/testify/assert"
)

var userInsertTests = []Contact{
	{FirstName: "Patrick", LastName: "Lecuyer", Email: "patate@mailinator.com"},
	{FirstName: "Patrick", LastName: "Lécuyer", Email: "patate2@mailinator.com"},
	{FirstName: "Marie-France", LastName: "Lemire", Email: "patate3@mailinator.com"},
}

func TestUserInsert(t *testing.T) {
	// config.LoadConfig()

	if err := config.LoadConfig(); err != nil {
		t.Skip("Couldn't load config, skipping test: ", err)
	}
	config.AppConfig.Datastore.GetDBSession().DropDatabase()
	defer config.AppConfig.Datastore.GetDBSession().DropDatabase()

	for _, u := range userInsertTests {
		if err := u.Insert(); err != nil {
			t.Error("Couldnt insert user: ", err)
		}

		r, err := ContactByEmail(u.Email)
		if err != nil {
			t.Error("Cannot find inserted user", err)
		}
		// Asserts that the user is the same
		assert.True(t, compare(*r, u))

		if err := u.Insert(); err != ErrContactAlreadyExists {
			t.Error("Didn't detect duplicate user: ", u, err)
		}

	}
}

var userUpdateTests = []struct {
	initial Contact
	changed Contact
	expect  bool
}{
	{
		Contact{FirstName: "Patrick", LastName: "Lecuyer", Email: "patate1@mailinator.com"},
		Contact{FirstName: "Vincent", LastName: "Lecuyer", Email: "patate1@mailinator.com"},

		true,
	}, {
		Contact{FirstName: "Patrick", LastName: "Lecuyer", Email: "patate2@mailinator.com"},
		Contact{FirstName: "Patrick", LastName: "Lécuyer", Email: "kapou@mailinator.com"},
		true,
	}, {
		Contact{FirstName: "Patrick", LastName: "Lecuyer", Email: "patate3@mailinator.com"},
		Contact{FirstName: "Patrick", LastName: "Lécuyer", Email: "patate3@mailinator.com"},
		true,
	},
}

func TestUserUpdate(t *testing.T) {
	config.LoadConfig()

	if err := config.LoadConfig(); err != nil {
		t.Skip("Couldn't load config, skipping test: ", err)
	}

	config.AppConfig.Datastore.GetDBSession().DropDatabase()
	defer config.AppConfig.Datastore.GetDBSession().DropDatabase()

	for _, i := range userUpdateTests {

		// Inserts the user
		assert.NoError(t, i.initial.Insert())

		// Validate that the user exists
		u, err := ContactByEmail(i.initial.Email)
		assert.NoError(t, err)

		// Validate that the user is the right one
		assert.True(t, compare(i.initial, *u))

		err = i.initial.Update(i.changed)

		if i.expect == true {
			assert.NoError(t, err)
			u, err = ContactByEmail(i.changed.Email)
			assert.NoError(t, err)
			assert.True(t, compare(i.changed, *u))
		} else {
			assert.Error(t, err, i.initial.Email, i.changed.Email)
		}

	}
}

var userCompareTests = []struct {
	a      Contact
	b      Contact
	expect bool
}{
	{
		Contact{FirstName: "Patrick", LastName: "Lecuyer", Email: "patate2@mailinator.com"},
		Contact{FirstName: "Patrick", LastName: "Lecuyer", Email: "patate2@mailinator.com"},
		true,
	},
	{
		Contact{FirstName: "Patrick", LastName: "Lecuyer", Email: "patate2@mailinator.com"},
		Contact{FirstName: "Patrick", LastName: "Lécuyer", Email: "patate2@mailinator.com"},
		false,
	},
	{
		Contact{FirstName: "Patrick", LastName: "Lecuyer", Email: "patate2@mailinator.com"},
		Contact{FirstName: "Patrick", LastName: "Patate", Email: "patate2@mailinator.com"},
		false,
	},
	{
		Contact{FirstName: "Patrick", LastName: "Lecuyer", Email: "patate2@mailinator.com"},
		Contact{FirstName: "Vincent", LastName: "Patate", Email: "patate2@mailinator.com"},
		false,
	},
	{
		Contact{FirstName: "Patrick", LastName: "Lecuyer", Email: "patate2@mailinator.com"},
		Contact{FirstName: "Patrick", LastName: "Lecuyer", Email: "patate@mailinator.com"},
		false,
	},
}

func TestUserCompare(t *testing.T) {

	for _, i := range userCompareTests {
		assert.Equal(t, i.expect, compare(i.a, i.b))
	}
}
