package repository

import "github.com/Yasinqurni/be-project/src/app/inventory/model"

type InventoryRepository interface {
	Create(inventory *model.Inventory) error
	Update(id uint, inventory *model.Inventory) error
	Get(id uint) (*model.Inventory, error)
	List() (*[]model.Inventory, error)
	Delete(id uint) error
}
