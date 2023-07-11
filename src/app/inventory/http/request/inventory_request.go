package request

type InventoryRequest struct {
	UserID    uint   `json:"user_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
	SkuNumber int    `json:"sku_number" validate:"required"`
	Notes     string `json:"notes" validate:"required"`
}
