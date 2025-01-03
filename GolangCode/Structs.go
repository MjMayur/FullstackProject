package main

type ResponseStruct struct {
	Code    int    `json:"code"`
	Massage string `json:"massage"`
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
