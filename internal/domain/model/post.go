package model

type Post struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" gorm:"size:255;not null"`
	Content  string `json:"content" gorm:"type:text;not null"`
	AuthorID uint
	Author   User `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TimestampTracking
}
