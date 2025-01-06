package main

type ResponseStruct struct {
	Code    int    `json:"code"`
	Massage string `json:"massage"`
}

type ResponseStructEntity struct {
	Code    int
	Massage string
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddUserRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type AddUserEntity struct {
	Name    string
	Email   string
	Phone   string
	Message string
}

// Define the userEntity struct
type UserEntity struct {
	Name     string
	Email    string
	Password string
}

type User struct {
	Name  string
	Email string
	Token string
}

type SuccessResponse struct {
	StatusCode int         `json:"code"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserListRes struct {
	ID      int    `json:"ID"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}
type ListResModal struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	Email   string `db:"email"`
	Phone   string `db:"phone"`
	Message string `db:"message"`
}
type ListResEntity struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Message string
}
