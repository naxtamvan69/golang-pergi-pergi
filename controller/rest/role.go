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

type RoleRestAPI interface {
	AddRole(w http.ResponseWriter, r *http.Request)
	UpdateRole(w http.ResponseWriter, r *http.Request)
	DeleteRole(w http.ResponseWriter, r *http.Request)
	GetRoles(w http.ResponseWriter, r *http.Request)
}

type roleRestAPI struct {
	roleService service.RoleService
}

func NewRoleRestAPI(roleService service.RoleService) *roleRestAPI {
	return &roleRestAPI{roleService}
}

func (ra *roleRestAPI) AddRole(w http.ResponseWriter, r *http.Request) {
	var roleRequest model.RoleRequest

	err := json.NewDecoder(r.Body).Decode(&roleRequest)
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

	role := model.Role{
		Role: roleRequest.Role,
	}

	newRole, err := ra.roleService.AddRoleService(r.Context(), role)
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success crate new role",
		Result:  newRole,
	})
}

func (ra *roleRestAPI) UpdateRole(w http.ResponseWriter, r *http.Request) {
	roleID := r.URL.Query().Get("role_id")
	if roleID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.FailedResponse{
			TimeStamp: time.Now(),
			Status:    http.StatusBadRequest,
			Error:     http.StatusText(http.StatusBadRequest),
			Message:   "role id is empty",
			Path:      r.RequestURI,
		})
		return
	}

	intRoleID, err := strconv.Atoi(roleID)
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

	var roleRequest model.RoleRequest

	err = json.NewDecoder(r.Body).Decode(&roleRequest)
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

	role := model.Role{
		ID:   intRoleID,
		Role: roleRequest.Role,
	}

	updateRole, err := ra.roleService.UpdateRoleService(r.Context(), role)
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
		Message: fmt.Sprintf("success update role with id = %d", intRoleID),
		Result:  updateRole,
	})
}

func (ra *roleRestAPI) DeleteRole(w http.ResponseWriter, r *http.Request) {
	roleID := r.URL.Query().Get("role_id")
	if roleID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.FailedResponse{
			TimeStamp: time.Now(),
			Status:    http.StatusBadRequest,
			Error:     http.StatusText(http.StatusBadRequest),
			Message:   "role id is empty",
			Path:      r.RequestURI,
		})
		return
	}

	intRoleID, err := strconv.Atoi(roleID)
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

	err = ra.roleService.DeleteRoleService(r.Context(), intRoleID)
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
		Message: fmt.Sprintf("success delete user with id = %d", intRoleID),
		Result: map[string]interface{}{
			"id": intRoleID,
		},
	})
}

func (ra *roleRestAPI) GetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := ra.roleService.GetRolesService(r.Context())
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
		Message: "success get all role",
		Result:  roles,
	})
}
