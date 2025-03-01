package controllers

import (
	"fmt"
	"harmonee/app/collections"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var StoreSession *sessions.CookieStore
var Validate *validator.Validate
var flashMessage = "flash-message"

type Server struct {
	Router *mux.Router
}

func (server *Server) Run(port string) {
	fmt.Printf("Listening to port %s", port)
	Validate = validator.New()
	StoreSession = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	log.Fatal(http.ListenAndServe(port, server.Router))
}

func GetPaginationLink(config *collections.AppConfig, params collections.PaginationParams) (collections.PaginationLink, error) {
	var links []collections.PageLink

	totalPage := int(math.Ceil(float64(params.TotalRow) / float64(params.PerPage)))

	for i := 1; i <= totalPage; i++ {
		links = append(links, collections.PageLink{
			Page:          i,
			Url:           fmt.Sprintf("%s/%s?page=%s", config.AppUrl, params.Path, fmt.Sprint(i)),
			IsCurrentPage: i == params.CurrentPage,
		})
	}

	var nextPage int
	var prevPage int

	prevPage = 1
	nextPage = totalPage

	if params.CurrentPage > 2 {
		prevPage = params.CurrentPage - 1
	}

	if params.CurrentPage < totalPage {
		nextPage = params.CurrentPage + 1
	}

	return collections.PaginationLink{
		CurrentPage: fmt.Sprintf("%s/%s?page=%s", config.AppUrl, params.Path, fmt.Sprint(params.CurrentPage)),
		NextPage:    fmt.Sprintf("%s/%s?page=%s", config.AppUrl, params.Path, fmt.Sprint(nextPage)),
		PrevPage:    fmt.Sprintf("%s/%s?page=%s", config.AppUrl, params.Path, fmt.Sprint(prevPage)),
		TotalRow:    params.TotalRow,
		TotalPage:   totalPage,
		Links:       links,
	}, nil
}

func SetFlashMessage(w http.ResponseWriter, r *http.Request, name string, value string) {
	session, err := StoreSession.Get(r, flashMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.AddFlash(value, name)
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func GetFlashMessage(w http.ResponseWriter, r *http.Request, name string) []string {
	session, err := StoreSession.Get(r, flashMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	var flashmsg []string
	if flashes := session.Flashes(name); len(flashes) > 0 {
		session.Save(r, w)
		for _, fl := range flashes {
			flashmsg = append(flashmsg, fl.(string))
		}

	}

	return flashmsg

}
