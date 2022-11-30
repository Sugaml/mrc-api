package controller

import (
	"encoding/json"
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
