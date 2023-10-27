package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	// hash password
	bytePw := []byte(c.Password)
	byteHashedPassword, err := bcrypt.GenerateFromPassword(bytePw, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	c.Password = string(byteHashedPassword)
	return
}
