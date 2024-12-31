package main

// ConvertUserEntity function
func ConvertUserEntity(request CreateUserRequest) UserEntity {
	return UserEntity{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

}
