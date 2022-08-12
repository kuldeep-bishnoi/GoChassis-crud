package usersservice

import (
	"log"
	"user_management/common"
	userrepo "user_management/repositories/users"
)

type UserService struct {
	UserRepo userrepo.UserRepoInterface
}

func (us *UserService) CreateUser(input common.CreateUserInput) common.Response {
	res, errcode, err := us.UserRepo.IsNameNotExists(input.Metadata["name"].(string))
	if err != nil {
		log.Println(err)
		return common.ResponseHandler(errcode, input.Language, 0, res)
	}
	res, errcode, err = us.UserRepo.Insert(input.Metadata)
	if err != nil {
		log.Println(err)
		return common.ResponseHandler(errcode, input.Language, 0, nil)
	}
	return common.ResponseHandler("701", input.Language, 1, res)
}
