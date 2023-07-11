package routes

import (
	"github.com/Yasinqurni/be-project/pkg/env"
	"github.com/Yasinqurni/be-project/src/app/inventory/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(group *gin.RouterGroup, db *gorm.DB, env *env.Config) {

	http.InventoryRoute(group, db, env)
}
