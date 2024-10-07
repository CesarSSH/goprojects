package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Api struct {
	Addr string
}

var users = []User{}

func (s *Api) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Prints status code 200
	w.WriteHeader(http.StatusOK)

	//Encode users slice to json
	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (s *Api) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	//Decode the body from request
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Create a new user
	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	//Insert the user and manage the errors
	if err = InserUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Prints status code 201
	w.WriteHeader(http.StatusCreated)

	message := map[string]string{"message": "User created successfully"}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func InserUser(u User) error {
	//Input validation
	if u.FirstName == "" {
		return errors.New("firstname is required")
	}
	if u.LastName == "" {
		return errors.New("lastname is required")
	}
	//Storage validation
	for _, user := range users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return errors.New("user already exists")
		}
	}
	users = append(users, u)

	return nil

}
