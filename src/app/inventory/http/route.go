package http

import (
	"github.com/Yasinqurni/be-project/pkg/env"
	"github.com/Yasinqurni/be-project/src/app/inventory"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InventoryRoute(group *gin.RouterGroup, db *gorm.DB, env *env.Config) {

	inventoryHandler, _, _ := inventory.InventoryInit(db, env)

	userGroup := group.Group("/inventory")

	userGroup.POST("/", inventoryHandler.Create)
	userGroup.PUT("/:id", inventoryHandler.Update)
	userGroup.DELETE("/:id", inventoryHandler.Delete)
	userGroup.GET("/", inventoryHandler.List)
	userGroup.GET("/:id", inventoryHandler.Detail)
}
