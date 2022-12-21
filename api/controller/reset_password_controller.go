package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sugam-project/api/models"
	"sugam-project/api/repository"
	"sugam-project/api/responses"
	"sugam-project/api/utils/mailer"
	"sugam-project/api/utils/security"
)

var resetPwd = repository.NewResetPassword()

func (server *Server) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer r.Body.Close()
	resetPasswordRequest := models.ResetPasswordRequest{}
	err = json.Unmarshal(body, &resetPasswordRequest)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user, err := urepo.FindbyUsername(server.DB, resetPasswordRequest.Email)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	resetPassword := models.ResetPassword{}
	repository.Prepare(&resetPassword)
	token := security.TokenHash(user.Email)
	resetPassword.Email = user.Email
	resetPassword.Token = token

	resetDetails, err := resetPwd.SaveDatails(server.DB, &resetPassword)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(resetDetails)
	mailer.SendResetPassword(resetDetails.Email, resetDetails.Token)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, map[string]interface{}{
		"Data":    resetDetails,
		"Success": 1,
	})
}

func (server *Server) ResetPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	requestBody := map[string]string{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	resetPassword, err := resetPwd.FindByToken(server.DB, requestBody["token"])
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	if resetPassword.DeletedAt != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("reset password link has been expired"))
		return
	}
	if requestBody["new_password"] == "" || requestBody["retype_password"] == "" {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("please ensure both field are entered"))
		return
	}
	if requestBody["new_password"] != "" && requestBody["retype_password"] != "" {
		if len(requestBody["new_password"]) < 6 || len(requestBody["retype_password"]) < 6 {
			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("password should be atleast 6 characters"))
			return
		}
		if requestBody["new_password"] != requestBody["retype_password"] {
			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("passwords provided do not match"))
			return
		}
		user.Password = requestBody["new_password"]
		user.Email = resetPassword.Email
		err = urepo.UpdatePassword(server.DB, &user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("cannot save, Please try again later"))
			return
		}
		_, err = resetPwd.DeleteDetails(server.DB, resetPassword)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, errors.New("cannot delete record, Please try again later"))
			return
		}
		responses.JSON(w, http.StatusOK, &map[string]string{
			"message": "Password reset successful",
		})
		return
	}
}
