package models

import "time"

type Product struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Title     string    `json:"title" gorm:"type: varchar(255)"`
	Price     int       `json:"price" gorm:"type: int"`
	Desc      string    `json:"desc" gorm:"type: varchar(255)"`
	Stock     int       `json:"stock" gorm:"type: int"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
