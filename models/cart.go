package models

import "time"

type Cart struct {
	ID            int                `json:"id" gorm:"primary_key:auto_increment"`
	Qty           int                `json:"qty"`
	SubTotal      int                `json:"subtotal"`
	UserID        int                `json:"user_id"`
	User          UserProfileRel     `json:"user"`
	ProductID     int                `json:"prouduct_id"`
	Product       ProductCartRel     `json:"product"`
	TransactionID *int               `json:"transaction_id"`
	Transaction   TransactionCartRel `json:"-"`
	CreatedAt     time.Time          `json:"-"`
	UpdatedAt     time.Time          `json:"updated_at"`
}
