package collections

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AppConfig struct {
	AppName string
	AppEnv  string
	AppUrl  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type PageLink struct {
	Page          int
	Url           string
	IsCurrentPage bool
}

type PaginationLink struct {
	CurrentPage string
	NextPage    string
	PrevPage    string
	TotalRow    int64
	TotalPage   int
	Links       []PageLink
}

type PaginationParams struct {
	Path        string
	TotalRow    int64
	PerPage     int
	CurrentPage int
}

func (s Server) InitializeApp(server Server) {
	panic("unimplemented")
}
