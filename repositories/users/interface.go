package userrepo

type UserRepoInterface interface {
	Insert(data map[string]interface{}) (map[string]interface{}, string, error)
	IsNameNotExists(name string) (map[string]interface{}, string, error)
}
