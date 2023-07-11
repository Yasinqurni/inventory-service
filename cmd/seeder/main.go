package main

import (
	"github.com/Yasinqurni/be-project/database/seeder"
	"github.com/Yasinqurni/be-project/pkg/database"
	"github.com/Yasinqurni/be-project/pkg/env"
)

func main() {

	config, _ := env.LoadConfig()
	db := database.InitDB(config)

	seeder.InventorySeed(db)
}
