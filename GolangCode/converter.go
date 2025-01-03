package main

// ConvertUserEntity function
func ConvertUserEntity(request CreateUserRequest) UserEntity {
	return UserEntity{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

}

func ConvertAddUserEntity(request AddUserRequest) AddUserEntity {
	return AddUserEntity{
		Name:    request.Name,
		Email:   request.Email,
		Phone:   request.Phone,
		Message: request.Message,
	}
}
