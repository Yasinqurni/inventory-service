package userRepository

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Yasinqurni/be-project/src/app/user/model"
)

type UserRepository interface {
	GetByID(id uint) (*model.User, error)
	GetByIDs(ids string) (*[]model.User, error)
}

type userRepositoryImpl struct {
	URL string
}

func NewUserRepositoryImpl(url string) UserRepository {
	return &userRepositoryImpl{
		URL: url,
	}
}

func (r *userRepositoryImpl) GetByID(id uint) (*model.User, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	fixURL := fmt.Sprintf("%s/%d", r.URL, id)
	fmt.Println(fixURL)
	resp, err := client.Get(fixURL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var data model.UserResponse
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	result := data.Data

	return &result, nil
}

func (r *userRepositoryImpl) GetByIDs(ids string) (*[]model.User, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	fixURL := fmt.Sprintf("%s/ids?ids=%s", r.URL, ids)
	fmt.Println(fixURL)
	resp, err := client.Get(fixURL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var data model.MultiUserResponse
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	result := data.Data

	return &result, nil
}
