package model

import (
	"database/sql"
	"time"

	userModel "github.com/Yasinqurni/be-project/src/app/user/model"
)

type Inventory struct {
	ID        uint
	UserID    uint
	Name      string
	Quantity  int
	SkuNumber int
	Notes     string
	User      userModel.User `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
