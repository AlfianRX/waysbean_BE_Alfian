package models

import "time"

type Cart struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Qty       int       `json:"qty" gorm:"type: int"`
	SubTotal  int       `json:"subtotal" gorm:"type: int"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"updated_at"`
}
