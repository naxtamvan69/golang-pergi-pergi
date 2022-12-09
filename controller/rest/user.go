package rest

import (
	"encoding/json"
	"net/http"
	"pergipergi/model"
	"pergipergi/service"
	"time"
)

type UserRestAPI interface {
	AddUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type userRestAPI struct {
	userService service.UserService
}

func NewUserRestAPI(userService service.UserService) *userRestAPI {
	return &userRestAPI{userService}
}

func (u *userRestAPI) AddUser(w http.ResponseWriter, r *http.Request) {
	var userRequest model.UserRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.FailedResponse{
			TimeStamp: time.Now(),
			Status:    http.StatusBadRequest,
			Error:     http.StatusText(http.StatusBadRequest),
			Message:   err.Error(),
			Path:      r.RequestURI,
		})
		return
	}

	user := model.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
		Name:     userRequest.Name,
		Sex:      userRequest.Sex,
		Age:      userRequest.Age,
	}

	newUser, err := u.userService.AddUserService(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.FailedResponse{
			TimeStamp: time.Now(),
			Status:    http.StatusInternalServerError,
			Error:     http.StatusText(http.StatusInternalServerError),
			Message:   err.Error(),
			Path:      r.RequestURI,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success crate new user",
		Result:  newUser,
	})
}

func (u *userRestAPI) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *userRestAPI) DeleteUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *userRestAPI) GetUsers(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
