package author

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
	authorsURL = "/authors"
	authorURL  = "/authors/:uuid"
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
	router.HandlerFunc(http.MethodGet, authorsURL, apperror.Middleware(h.GetAuthorsList))
	router.HandlerFunc(http.MethodPost, authorsURL, apperror.Middleware(h.CreateAuthor))
	router.HandlerFunc(http.MethodGet, authorURL, apperror.Middleware(h.GetAuthorByUUID))

}

func (h *handler) GetAuthorsList(w http.ResponseWriter, r *http.Request) error {

	//w.Write([]byte("Getting a list of authors"))
	//w.WriteHeader(200)

	return apperror.ErrNotFound
}

func (h *handler) GetAuthorByUUID(w http.ResponseWriter, r *http.Request) error {
	//w.Write([]byte("Getting author by uuid"))
	//w.WriteHeader(200)

	return apperror.NewAppError(nil, "test", "test", "t13")
}

func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request) error {
	//w.Write([]byte("Creating new author"))
	//w.WriteHeader(201)

	return fmt.Errorf("this is API error")
}

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
