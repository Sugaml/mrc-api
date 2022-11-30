package controller

import (
	"net/http"
	"sugam-project/api/repository"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	DB     *gorm.DB    //gorm db
	Router *mux.Router //router
	MRepo  interface{} //mockrepo for testing
}

func NewServer(db *gorm.DB) (*Server, error) {
	server := &Server{
		DB:     db,
		Router: mux.NewRouter(),
	}
	repository.Migrate(&repository.Repository{DB: server.DB})
	server.initializeRoutes()
	return server, nil
}

func (server *Server) Run(addr string) {
	log.Info("Listening to port http://localhost:", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
