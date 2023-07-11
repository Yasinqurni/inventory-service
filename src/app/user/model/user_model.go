package model

type UserResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    User   `json:"data,omitempty"`
}

type MultiUserResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    []User `json:"data,omitempty"`
}
type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}
