package usersservice

import "user_management/common"

type UsersserviceInterface interface {
	CreateUser(input common.CreateUserInput) common.Response
	GetAllUsers(input common.GetAllUsersInput) common.Response
	UpdateUserProfile(input common.UpdateUserProfileInput) common.Response
	DeleteUser(input common.DeleteUserInput) common.Response
	GetUserProfile(input common.GetUserProfileInput) common.Response
}
