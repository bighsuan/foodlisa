package models

import time "time"

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StoreID   int       `json:"storeId"`
	Sequence  int       `json:"sequence"`
	Name      string    `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
