package initproject

import (
	"fmt"

	"github.com/Yasinqurni/be-project/pkg/database"
	"github.com/Yasinqurni/be-project/pkg/env"
	"github.com/Yasinqurni/be-project/src/app/inventory/model"
	"github.com/Yasinqurni/be-project/src/routes"
	"github.com/gin-gonic/gin"
)

func InitProject() {

	config, _ := env.LoadConfig()

	db := database.InitDB(config)
	db.AutoMigrate(&model.Inventory{})

	r := gin.Default()

	v1 := r.Group("/api/v1")
	routes.Router(v1, db, config)

	r.Run(fmt.Sprintf(":%s", config.Port))
}
