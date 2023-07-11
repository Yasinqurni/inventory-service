package middleware

import (
	"github.com/Yasinqurni/be-project/src/app/inventory/http/request"
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(inventory request.InventoryRequest) bool {
	validate := validator.New()
	if err := validate.Struct(inventory); err != nil {
		return false
	}
	return true
}
