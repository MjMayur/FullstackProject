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
func ConvertUserModalToUserEntity(request ListResModal) ListResEntity {
	return ListResEntity{
		ID:      request.Id,
		Name:    request.Name,
		Email:   request.Email,
		Phone:   request.Phone,
		Message: request.Message,
	}
}

func ConvertUserEntityToUserJson(entity ListResEntity) UserListRes {
	return UserListRes{
		ID:      entity.ID,
		Name:    entity.Name,
		Email:   entity.Email,
		Phone:   entity.Phone,
		Message: entity.Message,
	}
}
