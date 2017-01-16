package handlers

import (
	"net/http"

	validator "gopkg.in/validator.v2"

	"github.com/gin-gonic/gin"

	"github.com/patricklecuyer/cr460-lab1/models"
)

// CreateContact handler
func CreateContact(c *gin.Context) {
	var cr = contactRequest{}

	if err := c.BindJSON(&cr); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate(cr); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	u := cr.Contact()
	if err := u.Insert(); err != nil {
		if err == models.ErrContactAlreadyExists {
			c.String(http.StatusConflict, "Contact already exists")
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	}

	c.Status(http.StatusCreated)

}

//ListContacts lists all contacts
func ListContacts(c *gin.Context) {
	contacts, err := models.GetAllContacts()

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, contacts)
}

//GetContact gets a contact
func GetContact(c *gin.Context) {
	email := c.Param("email")

	u, err := models.ContactByEmail(email)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, u)
}

// UpdateContact updates a Contact
func UpdateContact(c *gin.Context) {

	var cr = contactRequest{}

	if err := c.BindJSON(&cr); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate(cr); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	u := cr.Contact()

	f, err := models.ContactByEmail(u.Email)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	if f == nil {
		c.String(http.StatusNotFound, "Contact not found")
	}
	f.Update(u)
}
