package database

import (
	"fmt"

	"github.com/Yasinqurni/be-project/pkg/env"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c *env.Config) *gorm.DB {

	godotenv.Load(".env")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		c.DbHost, c.DbUsername, c.DbPassword, c.DbDatabase, c.DbPort, c.Sslmode, c.Timezone)
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DbUsername, c.DbPassword, c.DbHost, c.DbPort, c.DbDatabase)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
