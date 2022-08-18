package userrepo

type UserRepoInterface interface {
	Insert(data map[string]interface{}) (map[string]interface{}, string, error)
	IsNameNotExists(name string) (map[string]interface{}, string, error)
	GetAllUsersInput(filters string, page string, limit string) ([]map[string]interface{}, string, error, int)
	Delete(id string) (map[string]interface{}, string, error)
	Getbyid(id string) (map[string]interface{}, string, error)
	Update(id string, data map[string]interface{}) (map[string]interface{}, string, error)
}
