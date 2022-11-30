package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sugam-project/api/models"
	"sugam-project/api/repository"
	"sugam-project/api/responses"
)

var srepo = repository.NewStudentRepo()

func (server *Server) StudentInfo(w http.ResponseWriter, r *http.Request) {
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
	course, err := srepo.SaveStudent(server.DB, data)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, course)
}
