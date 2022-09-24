package params

import "time"

type OrderParams struct {
	CustomerName string       `json:"customer_name" binding:"required"`
	OrderedAt    time.Time    `json:"ordered_at" binding:"required"`
	Items        []ItemParams `json:"items" binding:"required"`
}
