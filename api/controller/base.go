package controller

import (
	"net/http"
	"os"
	"sugam-project/api/middleware"
	"sugam-project/api/repository"
	"sugam-project/api/utils/storage"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	DB            *gorm.DB    //gorm db
	Router        *mux.Router //router
	MRepo         interface{} //mockrepo for testing
	StorageClient storage.IStorage
}

func NewServer(db *gorm.DB) (*Server, error) {
	server := &Server{
		DB:     db,
		Router: mux.NewRouter(),
	}
	middleware.DB = server.DB
	repository.Migrate(&repository.Repository{DB: server.DB})
	server.initializeRoutes()
	stoarageClient, err := storage.NewStorage(os.Getenv("STORAGE_TYPE"))
	if err != nil {
		return nil, err
	}
	server.StorageClient = stoarageClient
	return server, nil
}

func (server *Server) Run(addr string) {
	log.Info("Listening to port http://localhost:", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
