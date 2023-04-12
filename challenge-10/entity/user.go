package entity

import (
	"challange-10/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email"`
	Password string    `gorm:"not null;" json:"-" `
	IsAdmin  bool      `gorm:"not null;default:0" json:"admin" form:"admin"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

type UserRequest struct {
	FullName string `form:"full_name" valid:"required"`
	Email    string `form:"email" valid:"required-Your email is required,email-Invalid email format"`
	Password string `form:"password" valid:"required,minstringlength(8)-Password has to have a minimum length of 8 characters"`
	IsAdmin  bool   `json:"admin" form:"admin"`
}

type UserLogin struct {
	Email    string `form:"email" valid:"required-Your email is required,email-Invalid email format"`
	Password string `form:"password" valid:"required,minstringlength(8)-Password has to have a minimum length of 8 characters"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
