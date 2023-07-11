package response

import (
	userModel "github.com/Yasinqurni/be-project/src/app/user/model"
)

type InventoryResponse struct {
	UserID    uint           `json:"user_id"`
	Name      string         `json:"name"`
	Quantity  int            `json:"quantity"`
	SkuNumber int            `json:"sku_number"`
	Notes     string         `json:"notes"`
	User      userModel.User `json:"user"`
}
