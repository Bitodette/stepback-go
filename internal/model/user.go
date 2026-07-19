package model

import "time"

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Email        string `gorm:"uniqueIndex;size:255;not null" json:"email"`
	PasswordHash string `gorm:"column:password_hash;size:255;not null" json:"-"`
	Name         string `gorm:"size:255;not null" json:"name"`

	PhoneNumber string `gorm:"column:phone_number;size:20" json:"phone_number"`

	// "customer", "admin", "staff"
	Role string `gorm:"size:20;default:customer" json:"role"`

	IsVerified bool `gorm:"column:is_verified;default:false" json:"is_verified"`

	// json:"-" = gak pernah dikirim ke client
	VerificationToken string    `gorm:"column:verification_token;size:255" json:"-"`
	ResetToken        string    `gorm:"column:reset_token;size:255" json:"-"`
	ResetTokenExpiry  time.Time `gorm:"column:reset_token_expires_at" json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
