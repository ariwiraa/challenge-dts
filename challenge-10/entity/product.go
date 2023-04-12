package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `gorm:"not null" json:"title" form:"title" valid:"required"`
	Description string `gorm:"not null" json:"description" form:"description" valid:"required"`
	UserId      uint
	User        *User
}

type ProductRequest struct {
	Title       string `json:"title" form:"title" valid:"required"`
	Description string `json:"description" form:"description" valid:"required"`
	UserId      uint
	User        *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
