package models

import "time"

type Profile struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	Phone     string    `json:"phone" gorm:"type: varchar(255)"`
	Address   string    `json:"address" gorm:"type: text"`
	PostCode  string    `json:"post_code" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
