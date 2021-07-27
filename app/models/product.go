package models

import time "time"

type Product struct {
	ID         uint      `json:"id" gorm:"primaryKey;auto_increase"`
	StoreID    int       `json:"storeId"`
	CategoryID int       `json:"categoryId"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Comment    string    `json:"comment"`
	ImageUrl   string    `json:"imageUrl"`
	CreatedAt  time.Time `json:""`
	UpdatedAt  time.Time `json:""`
	//Store	   Store     `json:"store,omitempty" gorm:"foreignKey:StoreID;references:ID"`
	//Category   Category  `json:"category,omitempty" gorm:"foreignKey:CategoryId;references:ID"`
}
