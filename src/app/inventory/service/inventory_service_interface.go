package service

import (
	"github.com/Yasinqurni/be-project/src/app/inventory/http/request"
	"github.com/Yasinqurni/be-project/src/app/inventory/model"
)

type InventoryService interface {
	Create(inventory *request.InventoryRequest) error
	Update(id uint, inventoryRequest request.InventoryRequest) error
	Get(id uint) (*model.Inventory, error)
	Delete(id uint) error
	List() (*[]model.Inventory, error)
}
