package controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sugam-project/api/models"
	"sugam-project/api/repository"
	"sugam-project/api/responses"

	"github.com/gorilla/mux"
)

var sfrepo = repository.NewStudentFileRepo()

func (server *Server) StudentFileInfo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer r.Body.Close()
	data := &models.StudentFile{}
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
	path, err := server.StorageClient.UplodaImage(context.Background(), data.SEECharacter, "certificate")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	data.CertifiacteTranscript = path.SecureURL
	sfile, err := sfrepo.SaveStudentFile(server.DB, data)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, sfile)
}

func (server *Server) UpdateStudentFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["sid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	studentFile, err := sfrepo.FindStudenFilebyId(server.DB, uint(sid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, studentFile)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	studentFileUpdated, err := sfrepo.UpdateStudentFile(server.DB, studentFile, uint(sid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, studentFileUpdated)
}

func (server *Server) DeleteStudentFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	sfile, err := sfrepo.DeleteStudentFile(server.DB, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, sfile)
}

func (server *Server) GetStudentFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["sid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	studentFile, err := sfrepo.FindStudenFilebyId(server.DB, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, studentFile)
}
