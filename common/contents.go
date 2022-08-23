package common

import "net/http"

type CreateUserInput struct {
	Metadata map[string]interface{}
	Language string
	Status   http.ResponseWriter
}

type GetAllUsersInput struct {
	ID       string
	Language string
	Page     string
	Size     string
	Filters  string
	Sort     string
	Status   http.ResponseWriter
}

type GetUserProfileInput struct {
	ID       string
	Language string
	Status   http.ResponseWriter
}

type DeleteUserInput struct {
	ID       string
	Language string
	Status   http.ResponseWriter
}

type UpdateUserProfileInput struct {
	ID       string
	Metadata map[string]interface{}
	Language string
	Status   http.ResponseWriter
}

var ErrorMessage map[string]interface{}
