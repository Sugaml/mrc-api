package middleware

import (
	"errors"
	"net/http"
	"sugam-project/api/auth"
	"sugam-project/api/repository"
	"sugam-project/api/responses"

	"github.com/jinzhu/gorm"
)

var (
	DB    *gorm.DB
	urepo = repository.NewUserRepo()
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func SetAdminMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		}
		userId, err := auth.ExtractTokenID(r)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		user, err := urepo.FindbyId(DB, uint(userId))
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
		if !user.IsAdmin {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("you are not authorized to use this api"))
			return
		}
		next.ServeHTTP(w, r)
	}
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
