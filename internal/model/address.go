package model

import "time"

type Address struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `gorm:"column:user_id;not null" json:"user_id"`

	// "Home", "Office", dll
	Label         string `gorm:"size:50;not null" json:"label"`
	RecipientName string `gorm:"column:recipient_name;size:255;not null" json:"recipient_name"`
	PhoneNumber   string `gorm:"column:phone_number;size:20;not null" json:"phone_number"`
	AddressLine   string `gorm:"column:address_line;not null" json:"address_line"`
	City          string `gorm:"size:100;not null" json:"city"`
	Province      string `gorm:"size:100;not null" json:"province"`
	PostalCode    string `gorm:"column:postal_code;size:10;not null" json:"postal_code"`

	// only one can be default per user
	IsDefault bool `gorm:"column:is_default;default:false" json:"is_default"`

	// for biteship ongkir calculation nanti
	BiteshipAreaID string `gorm:"column:biteship_area_id;size:255" json:"biteship_area_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Address) TableName() string {
	return "addresses"
}
