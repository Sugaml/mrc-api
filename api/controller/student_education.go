package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sugam-project/api/models"
	"sugam-project/api/repository"
	"sugam-project/api/responses"

	"github.com/gorilla/mux"
)

var serepo = repository.NewStudentEducationRepo()

func (server *Server) StudentEducation(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer r.Body.Close()
	datas := []models.StudentEducation{}
	err = json.Unmarshal(body, &datas)
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
	studenEdu := []models.StudentEducation{}
	for _, data := range datas {
		sedu, err := serepo.SaveStudentEducation(server.DB, &data)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		studenEdu = append(studenEdu, *sedu)
	}

	responses.JSON(w, http.StatusCreated, studenEdu)
}

func (server *Server) GetStudentEducation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["sid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	studentEdu, err := serepo.FindStudenEducationDetail(server.DB, uint(sid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, studentEdu)
}

func (server *Server) UpdateStudentEducation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["sid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	studentEdu, err := serepo.FindStudenEducationbyId(server.DB, uint(sid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, studentEdu)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	studentEduUpdated, err := serepo.UpdateStudentEducation(server.DB, studentEdu, uint(sid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, studentEduUpdated)
}

func (server *Server) DeleteStudentEducation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	sedu, err := serepo.DeleteStudentEducation(server.DB, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, sedu)
}
