package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"user_management/common"
	"user_management/database"
	userrepo "user_management/repositories/users"

	"user_management/resource"
	usersservice "user_management/services/users"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func LoadErrors(errors []map[string]interface{}) {
	res := make(map[string]interface{})
	for _, err := range errors {
		res[res["error"].(string)] = err
	}
	common.ErrorMessage = res
}

func main() {
	bytes, err := ioutil.ReadFile("C:/Users/Nviera/go/src/user_management/server/conf/errcodes.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	errors := make([]map[string]interface{}, 0)
	json.Unmarshal(bytes, &errors)
	LoadErrors(errors)
	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter()
	//Routes
	client := database.Connect()
	repo := userrepo.UserRepo{Database: "kuldeep", Client: client}
	r := resource.Resource{UserService: &usersservice.UserService{UserRepo: &repo}}
	s.HandleFunc("/createProfile", r.CreateUser).Methods("POST")
	s.HandleFunc("/getAllUsers", r.GetAllUsers).Methods("GET")
	// s.HandleFunc("/getUserProfile/{id}", getUserProfile).Methods("GET")
	// s.HandleFunc("/updateProfile/{id}", updateProfile).Methods("PUT")
	// s.HandleFunc("/deleteProfile/{id}", deleteProfile).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})
	handler := c.Handler(route)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
