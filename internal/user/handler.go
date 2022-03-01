package user

import (
	"fmt"
	"net/http"
	"restful_go_project/internal/apperror"
	"restful_go_project/pkg/logging"

	"restful_go_project/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetUsersList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))

}

func (h *handler) GetUsersList(w http.ResponseWriter, r *http.Request) error {

	//w.Write([]byte("Getting a list of users"))
	//w.WriteHeader(200)

	return apperror.ErrNotFound
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	//w.Write([]byte("Getting user by uuid"))
	//w.WriteHeader(200)

	return apperror.NewAppError(nil, "test", "test", "t13")
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	//w.Write([]byte("Creating new user"))
	//w.WriteHeader(201)

	return fmt.Errorf("this is API error")
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Update user"))
	w.WriteHeader(204)

	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Update user partially"))
	w.WriteHeader(204)

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Delete user"))
	w.WriteHeader(204)

	return nil
}
