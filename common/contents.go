package common

type CreateUserInput struct {
	Metadata map[string]interface{}
	Language string
}

type GetAllUsersInput struct {
	ID, Language, Page, Size, Filters, Sort string
}

type GetUserProfileInput struct {
	ID       string
	Language string
}

type DeleteUserInput struct {
	ID       string
	Language string
}

type UpdateUserProfileInput struct {
	ID       string
	Metadata map[string]interface{}
	Language string
}

var ErrorMessage map[string]interface{}
