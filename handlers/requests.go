package handlers

import "github.com/patricklecuyer/cr460-lab1/models"

type contactRequest struct {
	FirstName   string `json:"firstname" binding:"required"`
	LastName    string `json:"lastname" binding:"required" `
	Email       string `json:"email"  binding:"required" `
	PhoneNumber string `json:"phonenumber" `
}

func (r *contactRequest) Contact() models.Contact {
	c := models.Contact{}

	c.FirstName = r.FirstName
	c.LastName = r.LastName
	c.Email = r.Email
	c.PhoneNumber = r.PhoneNumber
	return c
}
