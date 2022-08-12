package usersservice

import "user_management/common"

type UsersserviceInterface interface {
	CreateUser(input common.CreateUserInput) common.Response
}
