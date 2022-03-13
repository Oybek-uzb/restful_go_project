package author

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restful_go_project/internal/apperror"
	service2 "restful_go_project/internal/author/service"
	"restful_go_project/pkg/api/sort"
	"restful_go_project/pkg/logging"
	"strconv"
	"strings"

	"restful_go_project/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

const (
	authorsURL = "/authors"
	authorURL  = "/authors/:uuid"
)

type handler struct {
	service *service2.Service
	logger  *logging.Logger
}

func NewHandler(service *service2.Service, logger *logging.Logger) handlers.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, authorsURL, sort.Middleware(apperror.Middleware(h.GetList), "created_at", sort.ASC))

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")
	if name != "" {

	}

	age := r.URL.Query().Get("age")
	if age != "" {
		operator := "="
		value := age
		if strings.Index(age, ":") != -1 {
			split := strings.Split(age, ":")
			operator = split[0]
			value = split[1]
		}
		fmt.Printf("operator: %s, value: %s", operator, value)
	}

	isAlive := r.URL.Query().Get("is_alive")
	if isAlive != "" {
		_, err := strconv.ParseBool(isAlive)
		if err != nil {
			validationError := apperror.BadRequestError("filter params validation failed", "wrong bool value")
			validationError.WithParams(map[string]string{
				"is_alive": "this field should be boolean: true or false",
			})
			return validationError
		}
	}

	createdAt := r.URL.Query().Get("created_at")
	if createdAt != "" {
		if strings.Index(createdAt, ":") != -1 {
			// range
		} else {
			// single
		}
	}

	var sortOptions sort.Options
	if options, ok := r.Context().Value(sort.OptionsContextKey).(sort.Options); ok {
		sortOptions = options
	}
	all, err := h.service.GetAll(r.Context(), sortOptions)
	if err != nil {
		w.WriteHeader(400)
		return err
	}
	allBytes, err := json.Marshal(all)
	if err != nil {
		return err
	}
	w.WriteHeader(200)
	w.Write(allBytes)

	return nil
}

//func (h *handler) GetAuthorByUUID(w http.ResponseWriter, r *http.Request) error {
//	//w.Write([]byte("Getting author by uuid"))
//	//w.WriteHeader(200)
//
//	return apperror.NewAppError(nil, "test", "test", "t13")
//}
//
//func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request) error {
//	//w.Write([]byte("Creating new author"))
//	//w.WriteHeader(201)
//
//	return fmt.Errorf("this is API error")
//}

//func (h *handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) error {
//	w.Write([]byte("Update author"))
//	w.WriteHeader(204)
//
//	return nil
//}
//
//func (h *handler) PartiallyUpdateAuthor(w http.ResponseWriter, r *http.Request) error {
//	w.Write([]byte("Update author partially"))
//	w.WriteHeader(204)
//
//	return nil
//}
//
//func (h *handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) error {
//	w.Write([]byte("Delete author"))
//	w.WriteHeader(204)
//
//	return nil
//}
