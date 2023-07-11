package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Yasinqurni/be-project/src/app/inventory/http/request"
	"github.com/Yasinqurni/be-project/src/app/inventory/http/response"
	"github.com/Yasinqurni/be-project/src/app/inventory/middleware"
	"github.com/Yasinqurni/be-project/src/app/inventory/service"
	"github.com/gin-gonic/gin"
)

type inventoryHendlerImpl struct {
	inventoryService service.InventoryService
}

func NewInventoryHendlerImpl(inventoryService service.InventoryService) InventoryHendler {
	return &inventoryHendlerImpl{
		inventoryService: inventoryService,
	}
}

func (h *inventoryHendlerImpl) Create(c *gin.Context) {

	var inventory request.InventoryRequest

	if err := c.ShouldBindJSON(&inventory); err != nil {
		data := response.NewErrorResponse("cannot Bind JSON", err.Error())
		c.JSON(http.StatusBadRequest, data)
		return
	}

	requestValid := middleware.ValidateStruct(inventory)

	if !requestValid {
		describeError := response.DescribeError{
			UserID:    "required",
			Name:      "required",
			Quantity:  "required",
			SkuNumber: "required",
			Notes:     "required",
		}
		data := response.NewErrorResponseValidate(describeError, "")
		c.JSON(http.StatusBadRequest, data)
		return
	}

	err := h.inventoryService.Create(&inventory)
	if err != nil {
		fmt.Println(err)
		data := response.NewErrorResponse("cannot create inventory", err.Error())
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	data := response.NewResponse("success create inventory", nil)
	c.JSON(http.StatusCreated, data)
}

func (h *inventoryHendlerImpl) Update(c *gin.Context) {

	var request request.InventoryRequest

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		data := response.NewErrorResponse("please insert id", err.Error())
		c.JSON(http.StatusBadRequest, data)
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		data := response.NewErrorResponse("cannot Bind JSON", err.Error())
		c.JSON(http.StatusBadRequest, data)
		return
	}

	inventory, err := h.inventoryService.Get(uint(id))
	if err != nil {
		data := response.NewErrorResponse("cannot get inventory", err.Error())
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	if inventory.Name == "" {
		data := response.NewErrorResponse("inventory not found", "")
		c.JSON(http.StatusNotFound, data)
		return
	}

	err = h.inventoryService.Update(uint(id), request)
	if err != nil {
		data := response.NewErrorResponse("cannot update inventory", err.Error())
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	data := response.NewResponse("success update inventory", nil)
	c.JSON(http.StatusOK, data)
}

func (h *inventoryHendlerImpl) Delete(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		data := response.NewErrorResponse("please insert id", err.Error())
		c.JSON(http.StatusBadRequest, data)
		return
	}

	inventory, err := h.inventoryService.Get(uint(id))
	if err != nil {
		data := response.NewErrorResponse("cannot find inventory", err.Error())
		c.JSON(http.StatusInternalServerError, data)
		return
	}
	if inventory.Name == "" {
		data := response.NewErrorResponse("inventory not found", "")
		c.JSON(http.StatusBadRequest, data)
		return
	}

	err = h.inventoryService.Delete(uint(id))
	if err != nil {
		data := response.NewErrorResponse("cannot delete inventory", err.Error())
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	data := response.NewResponse("success delete data", nil)
	c.JSON(http.StatusOK, data)
}

func (h *inventoryHendlerImpl) List(c *gin.Context) {

	inventories, err := h.inventoryService.List()
	if err != nil {
		data := response.NewErrorResponse("cannot list inventory", err.Error())
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	var datas []response.InventoryResponse
	for _, inventory := range *inventories {
		datas = append(datas, response.InventoryResponse{
			UserID:    inventory.UserID,
			Name:      inventory.Name,
			Quantity:  inventory.Quantity,
			SkuNumber: inventory.SkuNumber,
			Notes:     inventory.Notes,
			User:      inventory.User,
		})
	}

	data := response.NewResponse("success list data", datas)
	c.JSON(http.StatusOK, data)
}

func (h *inventoryHendlerImpl) Detail(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		data := response.NewErrorResponse("please insert id", err.Error())
		c.JSON(http.StatusBadRequest, data)
		return
	}

	inventory, err := h.inventoryService.Get(uint(id))
	if err != nil {
		data := response.NewErrorResponse("cannot detail inventory", err.Error())
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	result := response.InventoryResponse{
		UserID:    inventory.UserID,
		Name:      inventory.Name,
		Quantity:  inventory.Quantity,
		SkuNumber: inventory.SkuNumber,
		Notes:     inventory.Notes,
		User:      inventory.User,
	}
	if inventory.Name == "" {
		data := response.NewErrorResponse("inventory not found", err.Error())
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := response.NewResponse("success detail data", result)
	c.JSON(http.StatusOK, data)
}
