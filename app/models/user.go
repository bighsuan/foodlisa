package models

import time "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;auto_increase"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
