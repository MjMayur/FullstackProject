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

// Define the userEntity struct
type UserEntity struct {
	Name     string
	Email    string
	Password string
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
