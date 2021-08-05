package models

import time "time"

type Product struct {
	ID         uint      `json:"id" gorm:"primaryKey;auto_increase"`
	UserId    int       `json:"userId"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Description string    `json:"description"`
	ImageUrl   string    `json:"imageUrl"`
	CreatedAt  time.Time `json:""`
	UpdatedAt  time.Time `json:""`
}
