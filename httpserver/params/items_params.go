package params

type ItemParams struct {
	ItemCode    int    `json:"item_code" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	OrderId     int    `json:"orders_id"`
	ItemId      int    `json:"item_id"`
}
