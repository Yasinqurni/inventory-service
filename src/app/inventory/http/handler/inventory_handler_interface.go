package handler

import (
	"github.com/gin-gonic/gin"
)

type InventoryHendler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
	Detail(c *gin.Context)
}
