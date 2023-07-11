package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Yasinqurni/be-project/src/app/inventory/http/request"
	"github.com/Yasinqurni/be-project/src/app/inventory/model"
	"github.com/Yasinqurni/be-project/src/app/inventory/repository"
	modelUser "github.com/Yasinqurni/be-project/src/app/user/model"
	userService "github.com/Yasinqurni/be-project/src/app/user/service"
)

type inventoryServiceImpl struct {
	inventoryRepository repository.InventoryRepository
	userService         userService.UserService
}

func NewInventoryServiceImpl(inventoryRepository repository.InventoryRepository, userService userService.UserService) InventoryService {
	return &inventoryServiceImpl{
		inventoryRepository: inventoryRepository,
		userService:         userService,
	}
}

func (s *inventoryServiceImpl) Create(inventory *request.InventoryRequest) error {

	userResponse, err := s.userService.GetByID(inventory.UserID)
	if err != nil {
		return err
	}
	if userResponse.ID == 0 {
		return errors.New("user not found")

	}
	data := model.Inventory{
		UserID:    inventory.UserID,
		Name:      inventory.Name,
		Quantity:  inventory.Quantity,
		SkuNumber: inventory.SkuNumber,
		Notes:     inventory.Notes,
	}
	err = s.inventoryRepository.Create(&data)
	if err != nil {
		return err
	}

	return nil
}

func (s *inventoryServiceImpl) Update(id uint, inventory request.InventoryRequest) error {

	if inventory.UserID != 0 {
		userResponse, err := s.userService.GetByID(inventory.UserID)
		if err != nil {
			return err
		}
		if userResponse.ID == 0 {
			return errors.New("user not found")
		}

	}

	data := model.Inventory{
		UserID:    inventory.UserID,
		Name:      inventory.Name,
		Quantity:  inventory.Quantity,
		SkuNumber: inventory.SkuNumber,
		Notes:     inventory.Notes,
	}
	return s.inventoryRepository.Update(id, &data)

}

func (s *inventoryServiceImpl) Get(id uint) (*model.Inventory, error) {

	inventory, err := s.inventoryRepository.Get(id)
	if err != nil {
		return nil, err
	}

	user, err := s.userService.GetByID(inventory.UserID)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	result := model.Inventory{
		UserID:    inventory.UserID,
		Name:      inventory.Name,
		Quantity:  inventory.Quantity,
		SkuNumber: inventory.SkuNumber,
		Notes:     inventory.Notes,
		User:      *user,
	}
	return &result, nil
}

func (s *inventoryServiceImpl) Delete(id uint) error {

	return s.inventoryRepository.Delete(id)

}

func (s *inventoryServiceImpl) List() (*[]model.Inventory, error) {

	datas, err := s.inventoryRepository.List()
	if err != nil {
		return nil, err
	}

	var ids []string
	userMap := make(map[uint]modelUser.User)

	for _, d := range *datas {
		ids = append(ids, strconv.Itoa(int(d.UserID)))
	}

	idString := strings.Join(ids, ",")

	users, err := s.userService.GetByIDs(idString)
	if err != nil {
		return nil, err
	}

	for _, user := range *users {
		userMap[user.ID] = user
	}

	var inventories []model.Inventory

	for _, data := range *datas {
		if user, ok := userMap[data.UserID]; ok {
			inventories = append(inventories, model.Inventory{
				ID:        data.ID,
				UserID:    data.UserID,
				Name:      data.Name,
				Quantity:  data.Quantity,
				SkuNumber: data.SkuNumber,
				Notes:     data.Notes,
				User:      user,
			})
		}
	}

	return &inventories, nil
}
