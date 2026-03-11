// Package common contains common entities and value objects used across the application.
package common

import "time"

type TimestampTracking struct {
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
