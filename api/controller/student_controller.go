package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sugam-project/api/auth"
	"sugam-project/api/models"
	"sugam-project/api/repository"
	"sugam-project/api/responses"

	"github.com/gorilla/mux"
)

var srepo = repository.NewStudentRepo()

func (server *Server) StudentInfo(w http.ResponseWriter, r *http.Request) {
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := &models.Student{}
	err = json.Unmarshal(body, data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// data.Prepare()
	// err = data.Validate()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }
	data.UserId = uid
	course, err := srepo.SaveStudent(server.DB, data)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, course)
}

func (server *Server) UpdateStudentInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["sid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	studentInfo, err := srepo.FindbyId(server.DB, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, studentInfo)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	studentUpdated, err := srepo.UpdateStudent(server.DB, studentInfo, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, studentUpdated)
}

func (server *Server) StudentGeneralInfo(w http.ResponseWriter, r *http.Request) {
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	student, err := srepo.FindbyUserId(server.DB, uint(uid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusCreated, student)
}

func (server *Server) StudentDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	student, err := srepo.FindbyId(server.DB, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusCreated, student)
}

func (server *Server) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["sid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	student, err := srepo.FindbyId(server.DB, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	_, err = srepo.DeleteStudent(server.DB, student.ID)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, "Deleted Successfully")
}
