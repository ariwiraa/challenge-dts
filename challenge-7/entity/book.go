package entity

type Book struct {
	Id     int    `gorm:"primaryKey" json:"id"`
	Title  string `gorm:"not null;type:varchar(100)" json:"title"`
	Author string `gorm:"not null;type:varchar(100)" json:"author"`
	Desc   string `gorm:"not null;type:ext" json:"desc"`
}

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}
