package seeder

import (
	"fmt"
	"log"

	"github.com/Yasinqurni/be-project/src/app/inventory/model"
	"gorm.io/gorm"
)

func InventorySeed(db *gorm.DB) {

	data := make([]model.Inventory, 3)

	for i := 0; i < 3; i++ {
		inventory := model.Inventory{
			UserID:    uint(i),
			Name:      fmt.Sprintf("item %d", i+1),
			Quantity:  int(20 + i),
			SkuNumber: int(2023 + i),
			Notes:     fmt.Sprintf("item ini adalah item%d", i+1),
		}
		data[i] = inventory
	}

	err := db.Create(&data).Error
	if err != nil {
		log.Fatal("Error ", err)
	}

}
