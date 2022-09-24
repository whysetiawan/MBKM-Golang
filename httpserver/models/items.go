package models

import (
	"time"
)

type Item struct {
	ItemId      uint      `gorm:"primaryKey" json:"item_id"`
	ItemCode    int       `json:"item_code"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderId     int       `json:"orders_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
