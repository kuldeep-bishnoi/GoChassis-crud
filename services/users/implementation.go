package usersservice

import (
	"log"
	"user_management/common"
	userrepo "user_management/repositories/users"
)

type UserService struct {
	UserRepo userrepo.UserRepoInterface
}

func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }

func (us *UserService) CreateUser(input common.CreateUserInput) common.Response {
	res, errcode, err := us.UserRepo.IsNameNotExists(input.Metadata["name"].(string))
	if err != nil {
		log.Println(err)
		return common.ResponseHandler(errcode, input.Language, 0, res, input.Status)
	}
	res, errcode, err = us.UserRepo.Insert(input.Metadata)
	if err != nil {
		log.Println(err)
		return common.ResponseHandler(errcode, input.Language, 0, nil, input.Status)
	}
	return common.ResponseHandler("701", input.Language, 1, res, input.Status)
}

func (us *UserService) GetAllUsers(input common.GetAllUsersInput) common.Response {
	res, errcode, err, count := us.UserRepo.GetAllUsersInput(input.Filters, input.Page, input.Size)
	if err != nil {
		log.Println(err)
		return common.ResponseHandler(errcode, input.Language, count, nil, input.Status)
	}
	return common.ResponseHandler("710", input.Language, count, res, input.Status)
}

func (us *UserService) DeleteUser(input common.DeleteUserInput) common.Response {

	res, errcode, err := us.UserRepo.Delete(input.ID)
	if err != nil {
		log.Println(err)
		return common.ResponseHandler(errcode, input.Language, 0, nil, input.Status)
	}
	return common.ResponseHandler("708", input.Language, 1, res, input.Status)
}

func (us *UserService) GetUserProfile(input common.GetUserProfileInput) common.Response {
	res, errcode, err := us.UserRepo.Getbyid(input.ID)
	if err != nil {
		log.Println(err)
		return common.ResponseHandler(errcode, input.Language, 0, nil, input.Status)
	}
	return common.ResponseHandler("710", input.Language, 1, res, input.Status)
}
func (us *UserService) UpdateUserProfile(input common.UpdateUserProfileInput) common.Response {
	res, errcode, err := us.UserRepo.Getbyid(input.ID)
	if err != nil {
		log.Println(err)
		return common.ResponseHandler(errcode, input.Language, 0, nil, input.Status)
	}
	name, nok := input.Metadata["name"].(string)
	if nok {
		if name != res["name"].(string) {
			res, errcode, err := us.UserRepo.IsNameNotExists(input.Metadata["name"].(string))
			if err != nil {
				log.Println(err)
				return common.ResponseHandler(errcode, input.Language, 0, res, input.Status)
			}
		}
	}
	res, errcode, err = us.UserRepo.Update(input.ID, input.Metadata)
	if err != nil {
		log.Println(err)
		return common.ResponseHandler(errcode, input.Language, 0, nil, input.Status)
	}
	return common.ResponseHandler("714", input.Language, 1, res, input.Status)
}
