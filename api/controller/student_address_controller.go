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

var arepo = repository.NewStudentAddressRepo()

func (server *Server) StudentAddress(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer r.Body.Close()
	data := &models.Address{}
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
	address, err := arepo.SaveStudentAddress(server.DB, data)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, address)
}

func (server *Server) UpdateStudentAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["sid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	studentAddress, err := arepo.FindStudentAddressbyId(server.DB, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, studentAddress)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	addressUpdated, err := arepo.UpdateStudentAddress(server.DB, studentAddress, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, addressUpdated)
}

func (server *Server) GetStudentAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["sid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	studentAddress, err := arepo.FindStudentAddressbyId(server.DB, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, studentAddress)
}
