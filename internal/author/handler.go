package author

import (
	"encoding/json"
	"net/http"
	"restful_go_project/internal/apperror"
	service2 "restful_go_project/internal/author/service"
	"restful_go_project/pkg/api/filter"
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
	router.HandlerFunc(http.MethodGet, authorsURL, filter.Middleware(sort.Middleware(apperror.Middleware(h.GetList), "created_at", sort.ASC), 10))

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {

	filterOptions := r.Context().Value(filter.OptionsContextKey).(filter.Options)

	name := r.URL.Query().Get("name")
	if name != "" {
		err := filterOptions.AddField("name", filter.OperatorLike, name, filter.DataTypeString)
		if err != nil {
			return err
		}
	}

	age := r.URL.Query().Get("age")
	if age != "" {
		operator := filter.OperatorEqual
		value := age
		if strings.Index(age, ":") != -1 {
			split := strings.Split(age, ":")
			operator = split[0]
			value = split[1]
		}
		err := filterOptions.AddField("age", operator, value, filter.DataTypeInt)
		if err != nil {
			return err
		}
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
		err = filterOptions.AddField("is_alive", filter.OperatorEqual, isAlive, filter.DataTypeBool)
		if err != nil {
			return err
		}
	}

	createdAt := r.URL.Query().Get("created_at")
	if createdAt != "" {
		var operator string
		if strings.Index(createdAt, ":") != -1 {
			// range
			operator = filter.OperatorBetween
		} else {
			// single
			operator = filter.OperatorEqual
		}
		err := filterOptions.AddField("created_at", operator, createdAt, filter.DataTypeDate)
		if err != nil {
			return err
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
