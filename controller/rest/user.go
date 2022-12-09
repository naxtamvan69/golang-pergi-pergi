package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pergipergi/model"
	"pergipergi/service"
	"strconv"
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
		Username:       userRequest.Username,
		Password:       userRequest.Password,
		Name:           userRequest.Name,
		Sex:            userRequest.Sex,
		Age:            userRequest.Age,
		TravelAgensiID: userRequest.TravelAgensiID,
		RoleID:         userRequest.TravelAgensiID,
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
	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.FailedResponse{
			TimeStamp: time.Now(),
			Status:    http.StatusBadRequest,
			Error:     http.StatusText(http.StatusBadRequest),
			Message:   "user id is empty",
			Path:      r.RequestURI,
		})
		return
	}

	_, err := strconv.Atoi(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.FailedResponse{
			TimeStamp: time.Now(),
			Status:    http.StatusBadRequest,
			Error:     http.StatusText(http.StatusBadRequest),
			Message:   "user id is not integer",
			Path:      r.RequestURI,
		})
		return
	}

	var userRequest model.UserRequest

	err = json.NewDecoder(r.Body).Decode(&userRequest)
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
		ID:             userRequest.ID,
		Username:       userRequest.Username,
		Password:       userRequest.Password,
		Name:           userRequest.Name,
		Sex:            userRequest.Sex,
		Age:            userRequest.Age,
		TravelAgensiID: userRequest.TravelAgensiID,
		RoleID:         userRequest.TravelAgensiID,
	}

	newUser, err := u.userService.UpdateUserService(r.Context(), user)
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

func (u *userRestAPI) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.FailedResponse{
			TimeStamp: time.Now(),
			Status:    http.StatusBadRequest,
			Error:     http.StatusText(http.StatusBadRequest),
			Message:   "user id is empty",
			Path:      r.RequestURI,
		})
		return
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.FailedResponse{
			TimeStamp: time.Now(),
			Status:    http.StatusBadRequest,
			Error:     http.StatusText(http.StatusBadRequest),
			Message:   "user id is not integer",
			Path:      r.RequestURI,
		})
		return
	}

	err = u.userService.DeleteUserService(r.Context(), intUserID)
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
		Status:  http.StatusOK,
		Message: fmt.Sprintf("success delete user with id = %d", intUserID),
		Result: map[string]interface{}{
			"id": intUserID,
		},
	})
}

func (u *userRestAPI) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.userService.GetUsersService(r.Context())
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
		Status:  http.StatusOK,
		Message: "success get list user",
		Result:  users,
	})
}
