package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sugam-project/api/auth"
	"sugam-project/api/models"
	"sugam-project/api/repository"
	"sugam-project/api/responses"

	"github.com/gorilla/mux"
)

var repo = repository.NewCourseRepo()

// Register godoc
// @Summary register to 01cloud
// @Description register to 01cloud
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body models.Course true "User register"
// @Param gitid query string true "Git Id"
// @Success 200 {object} models.Course
// @Router /course [post]
func (server *Server) CreateCourse(w http.ResponseWriter, r *http.Request) {
	err := server.CheckAdminAuthorization(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer r.Body.Close()
	data := &models.Course{}
	err = json.Unmarshal(body, data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = data.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	course, err := repo.SaveCourse(server.DB, data)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, course)
}

// ListAccounts godoc
// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {array}   models.Course
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /courses [get]
func (server *Server) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	course, err := repo.FindbyId(server.DB, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, course)
}

func (server *Server) GetCourses(w http.ResponseWriter, r *http.Request) {
	course, err := repo.FindAllCourse(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, course)
}

func (server *Server) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	course, err := repo.FindbyId(server.DB, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, course)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	courseUpdated, err := repo.UpdateCourse(server.DB, course, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, courseUpdated)
}

func (server *Server) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = repo.DeleteCourse(server.DB, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "Deleted successfully")
}

func (server *Server) CheckAuthorization(r *http.Request) error {
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		return err
	}
	user, err := urepo.FindbyId(server.DB, uid)
	if err != nil {
		return err
	}
	if !user.Active {
		return err
	}
	return nil
}

func (server *Server) CheckAdminAuthorization(r *http.Request) error {
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		return err
	}
	fmt.Println("user id ", uid)
	user, err := urepo.FindbyId(server.DB, uid)
	if err != nil {
		return err
	}
	if !user.Active && !user.IsAdmin {
		return err
	}
	return nil
}
