package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sugam-project/api/auth"
	"sugam-project/api/models"
	"sugam-project/api/repository"
	"sugam-project/api/responses"
	"sugam-project/api/utils/mailer"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var urepo = repository.NewUserRepo()

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	req := &models.UserRequest{}
	err = json.Unmarshal(body, req)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.NewUser(req)
	data.Prepare()
	err = data.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = urepo.FindbyUsername(server.DB, req.Email)
	if err == nil {
		responses.ERROR(w, http.StatusConflict, errors.New("already registered email"))
		return
	}
	user, err := urepo.Save(server.DB, data)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	_, err = srepo.SaveStudent(server.DB, models.NewStudent(user.ID, req))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	token, err := auth.CreateToken(user.ID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	err = mailer.SendVerifyEmail(user.Email, token)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// sender := mailer.NewGmailSender("Babulal", "tamangsugam09@gmail.com", "eokbjhvwagftsaca")
	// subject := "A test email"
	// content := `
	// <h1>Welcome to page</h1>
	// `
	// to := []string{data.Email}
	// attachFiles := []string{}
	// err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	// if err != nil {
	// 	fmt.Print(fmt.Errorf("error to send email", err))
	// 	return
	// }
	// fmt.Print("Success")
	responses.JSON(w, http.StatusCreated, user)
}

func (server *Server) GetUserByID(w http.ResponseWriter, r *http.Request) {
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user, err := urepo.FindbyId(server.DB, uint(uid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user, err := urepo.FindbyId(server.DB, uint(uid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := urepo.FindAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) UserEmailVerfy(w http.ResponseWriter, r *http.Request) {
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user, err := urepo.FindbyId(server.DB, uint(uid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	_, err = urepo.VerifyEmail(server.DB, user.ID, true)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user, err := urepo.FindbyId(server.DB, uint(uid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	courseUpdated, err := urepo.Update(server.DB, user, uint(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, courseUpdated)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = urepo.Delete(server.DB, uint(uid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "Deleted successfully")
}

func (server *Server) GetLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := &models.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = user.LoginValidate()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	data, err := urepo.FindbyUsername(server.DB, user.Username)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	if !data.EmailVerified {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("please verify your email"))
		return
	}
	log.Info("Login user id :: ", data.ID)
	token, err := auth.CreateToken(data.ID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	// sender := mailer.NewGmailSender("Babulal", "tamangsugam09@gmail.com", "eokbjhvwagftsaca")
	// subject := "A test email"
	// content := `
	// <h1>Welcome to page</h1>
	// `
	// to := []string{data.Email}
	// attachFiles := []string{}
	// err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	// if err != nil {
	// 	fmt.Print(fmt.Errorf("error to send email", err))
	// 	return
	// }
	// fmt.Print("Success")
	responses.JSON(w, http.StatusOK, map[string]interface{}{
		"token": token,
		"user":  data,
	})
}

func (server *Server) ActiveAndDeactiveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user, err := urepo.FindbyId(server.DB, uint(cid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	studentStatusRequest := &models.StudentStatusRequest{}
	err = json.Unmarshal(body, studentStatusRequest)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userUpdated, err := urepo.ActiveDeactiveUser(server.DB, user.ID, studentStatusRequest.Status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, userUpdated)
}
