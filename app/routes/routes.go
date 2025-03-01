package routes

import (
	"fmt"
	"harmonee/app/controllers"
	"harmonee/app/repositories"
	"harmonee/app/services"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes(server *controllers.Server) {
	fmt.Println("Initialize router...")
	server.Router = mux.NewRouter()

	server.Router.HandleFunc("/", controllers.Dashboard).Methods("GET")

	// SECTION ROUTE
	sectionRepo := repositories.NewSectionRepository()
	serviceSection := services.NewSectionService(&sectionRepo)
	sectionController := controllers.NewSectionController(&serviceSection)
	server.Router.HandleFunc("/sections", sectionController.FindAll).Methods("GET")
	server.Router.HandleFunc("/create-section", sectionController.Create).Methods("GET")
	server.Router.HandleFunc("/section/{id}", sectionController.Detail).Methods("GET")
	server.Router.HandleFunc("/section/add", sectionController.Store).Methods("POST")
	server.Router.HandleFunc("/section/update/{id}", sectionController.Update).Methods("POST")
	server.Router.HandleFunc("/section/delete/{id}", sectionController.Delete).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
