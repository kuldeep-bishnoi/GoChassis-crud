package resource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user_management/common"
	"user_management/handlers"
	us "user_management/services/users"

	"github.com/gorilla/mux"
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
	var parms = mux.Vars(r)
	id := parms["id"]
	page := parms["page"]
	size := parms["size"]
	filters := parms["filters"]
	var person = make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&person)
	ip := common.GetAllUsersInput{ID: id, Language: "en", Page: page, Size: size, Filters: filters}
	response := sr.UserService.GetAllUsers(ip)
	json.NewEncoder(w).Encode(response)
}

func (sr *Resource) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id := params["id"]
	input := common.DeleteUserInput{ID: id}
	var person = make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&person)
	ip := common.DeleteUserInput{ID: input.ID, Language: "en"}
	response := sr.UserService.DeleteUser(ip)
	json.NewEncoder(w).Encode(response)

}

func (sr *Resource) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id := params["id"]
	var person = make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&person)
	ip := common.GetUserProfileInput{ID: id, Language: "en"}
	response := sr.UserService.GetUserProfile(ip)
	json.NewEncoder(w).Encode(response)
}

func (sr *Resource) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id := params["id"]
	var person = make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&person)
	cpath := "C:/Users/Nviera/Desktop/schemaupdate.json"
	result, err := handlers.Validate(cpath, person)
	if err != nil {
		fmt.Println(err)
		common.ResponseHandler("", "en", 0, result)
		json.NewEncoder(w).Encode(result)
		return
	}
	ip := common.UpdateUserProfileInput{ID: id, Metadata: person, Language: "en"}
	response := sr.UserService.UpdateUserProfile(ip)
	json.NewEncoder(w).Encode(response)
}
