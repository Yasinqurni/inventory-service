package inventory

import (
	"github.com/Yasinqurni/be-project/pkg/env"
	handler "github.com/Yasinqurni/be-project/src/app/inventory/http/handler"
	"github.com/Yasinqurni/be-project/src/app/inventory/repository"
	service "github.com/Yasinqurni/be-project/src/app/inventory/service"
	userRepository "github.com/Yasinqurni/be-project/src/app/user/repository"
	userService "github.com/Yasinqurni/be-project/src/app/user/service"
	"gorm.io/gorm"
)

func InventoryInit(db *gorm.DB, env *env.Config) (handler.InventoryHendler, service.InventoryService, repository.InventoryRepository) {

	inventoryRepository := repository.NewInventoryRepositoryImpl(db)
	userRepository := userRepository.NewUserRepositoryImpl(env.Url)
	userService := userService.NewUserServiceImpl(userRepository)
	inventoryService := service.NewInventoryServiceImpl(inventoryRepository, userService)
	inventoryHandler := handler.NewInventoryHendlerImpl(inventoryService)

	return inventoryHandler, inventoryService, inventoryRepository
}
