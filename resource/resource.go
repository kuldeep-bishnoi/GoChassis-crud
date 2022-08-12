package resource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user_management/common"
	"user_management/handlers"
	us "user_management/services/users"
)

type Resource struct {
	UserService us.UsersserviceInterface
}

func (sr *Resource) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person = make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&person)
	cpath := "C:/Users/Nviera/Desktop/schemainsert.json"
	result, err := handlers.Validate(cpath, person)
	if err != nil {
		fmt.Println(err)
		common.ResponseHandler("", "en", 0, result)
		json.NewEncoder(w).Encode(result)
		return
	}
	ip := common.CreateUserInput{Metadata: person, Language: "en"}
	response := sr.UserService.CreateUser(ip)
	json.NewEncoder(w).Encode(response)
}

func (sr *Resource) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person = make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&person)

	ip := common.FetchAllUsersInput{ID: "", Filters: "", Language: "en", Page: "", Size: ""}
	response := sr.UserService.FetchAllUsersInput(ip)
	json.NewEncoder(w).Encode(response)
}
