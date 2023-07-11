package repository

import (
	"database/sql"
	"time"

	"github.com/Yasinqurni/be-project/src/app/inventory/model"
	"gorm.io/gorm"
)

type inventoryRepositoryImpl struct {
	db *gorm.DB
}

func NewInventoryRepositoryImpl(db *gorm.DB) InventoryRepository {
	return &inventoryRepositoryImpl{
		db: db,
	}
}

func (r *inventoryRepositoryImpl) Create(inventory *model.Inventory) error {
	err := r.db.Create(inventory).Error

	return err
}

func (r *inventoryRepositoryImpl) Update(id uint, inventory *model.Inventory) error {
	err := r.db.Where("id = ?", id).Updates(&inventory).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *inventoryRepositoryImpl) Get(id uint) (*model.Inventory, error) {
	var inventory model.Inventory
	err := r.db.Where("id = ?", id).Where("deleted_at = ?", "").Find(&inventory).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (r *inventoryRepositoryImpl) Delete(id uint) error {

	var inventory model.Inventory

	result := r.db.First(&inventory, id)
	if result.Error != nil {
		return result.Error
	}

	// Set DeletedAt dengan waktu saat ini
	now := time.Now()
	inventory.DeletedAt = sql.NullTime{Time: now, Valid: true}

	// Simpan perubahan ke database
	result = r.db.Save(&inventory)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *inventoryRepositoryImpl) List() (*[]model.Inventory, error) {
	var inventory []model.Inventory
	err := r.db.Find(&inventory).Where("deleted_at = ?", "").Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}
