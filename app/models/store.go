package models

import time "time"

type Store struct {
	ID                 int       `json:"id" gorm:"primaryKey;auto_increase"`
	UserID             int       `json:"userId"`
	Address            string    `json:"address"`
	Phone              string    `json:"phone"`
	ContractBeginTime  time.Time `json:"contractBeginTime"`
	ContractExpireTime time.Time `json:"contractExpireTime"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
