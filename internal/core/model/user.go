package model

type User struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Email          string `json:"email" gorm:"uniqueIndex;not null"`
	Firstname      string `json:"firstname" gorm:"size:255;not null"`
	Lastname       string `json:"lastname" gorm:"size:255;not null"`
	HashedPassword string `json:"hashedPassword" gorm:"size:255;not null"`
	RefreshToken   string `json:"refreshToken" gorm:"size:255"`
	IsActive       bool   `json:"isActive" gorm:"default:true"`
	TimestampTracking
}
