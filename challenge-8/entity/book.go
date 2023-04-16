package entity

import "time"

// Book represents the model for an book
type Book struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null;type:varchar(100)" json:"title"`
	Author    string    `gorm:"not null;type:varchar(100)" json:"author"`
	Desc      string    `gorm:"not null;type:text" json:"desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}
