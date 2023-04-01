package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sugam-project/api/auth"
	"sugam-project/api/middleware"
	"sugam-project/api/responses"

	_ "sugam-project/docs"

	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (server *Server) setJSON(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.HandleFunc(path, middleware.SetMiddlewareJSON(next)).Methods(method, "OPTIONS")
}

func (server *Server) setAdmin(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.setJSON(path, middleware.SetAdminMiddlewareAuthentication(next), method)
}

func (server *Server) initializeRoutes() {
	server.Router.Use(middleware.CORS)

	server.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	server.setJSON("/", server.WelcomePage, "GET")
	server.SetRoutes("/payment", "PAYMENT_SERVER_URL")
	server.SetRoutes("/uploads", "UPLOAD_SERVER_URL")

	server.setJSON("/course", server.CreateCourse, "POST")
	server.setJSON("/user/upload", server.handleFileupload, "POST")
	server.setJSON("/course/{id}", server.GetCourseByID, "GET")
	server.setJSON("/courses", server.GetCourses, "GET")
	server.setAdmin("/course/{id}", server.UpdateCourse, "PUT")
	server.setAdmin("/course/{id}", server.DeleteCourse, "DELETE")

	server.setJSON("/user", server.CreateUser, "POST")
	server.setJSON("/user/login", server.GetLogin, "POST")
	server.setJSON("/user", server.GetUserByID, "GET")
	server.setJSON("/user/{id}", server.GetUser, "GET")
	server.setJSON("/users", server.GetUsers, "GET")
	server.setJSON("/user/{id}", server.GetUser, "GET")
	server.setJSON("/user/{id}", server.UpdateUser, "PUT")
	server.setJSON("/user/email/verify", server.UserEmailVerfy, "PUT")
	server.setJSON("/user/{id}/active-deactive", server.ActiveAndDeactiveUser, "PUT")
	server.setJSON("/user/{id}", server.DeleteUser, "DELETE")
	server.setJSON("/user/forgot-password", server.ForgotPassword, "POST")
	server.setJSON("/user/reset-password", server.ResetPassword, "POST")

	server.setJSON("/student_info", server.StudentInfo, "POST")
	server.setJSON("/students", server.ListStudents, "GET")
	server.setJSON("/student_info/{id}", server.StudentDetail, "GET")
	server.setJSON("/student/general-info", server.StudentGeneralInfo, "GET")
	server.setJSON("/student/address", server.StudentAddress, "POST")
	server.setJSON("/student/education", server.StudentEducation, "POST")
	server.setJSON("/student/{sid}/status", server.UpdateStudentStatus, "PUT")
	server.setJSON("/student/file", server.StudentFileInfo, "POST")
	server.setJSON("/student/{sid}/address", server.GetStudentAddress, "GET")
	server.setJSON("/student/{sid}/education", server.GetStudentEducation, "GET")
	server.setJSON("/student/{sid}/document", server.GetStudentFile, "GET")

	server.setJSON("/student/{sid}", server.UpdateStudentInfo, "PUT")
	server.setJSON("/student/{sid}/address", server.UpdateStudentAddress, "PUT")
	server.setJSON("/student/{sid}/education", server.UpdateStudentEducation, "PUT")
	server.setJSON("/student/{sid}/document", server.UpdateStudentFile, "PUT")
}

func (server *Server) WelcomePage(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "welcome to mrc-api service")
}

func (server *Server) SetRoutes(path string, envValue string) {
	server.Router.PathPrefix(path).HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewReader(body))
		url := fmt.Sprintf("%s%s", os.Getenv(envValue), r.RequestURI)
		proxyReq, err := http.NewRequest(r.Method, url, bytes.NewReader(body))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadGateway)
			return
		}
		copyHeader(proxyReq.Header, r.Header)
		userId, _ := auth.ExtractTokenID(r)
		if userId != 0 {
			userGotten, _ := urepo.FindbyId(server.DB, userId)
			if userGotten != nil {
				logrus.Infof("proxy passed by :: %s :: %w", userGotten.ID, userGotten.IsAdmin)
			}
			proxyReq.Header.Set("x-user-id", fmt.Sprint(userId))
			if userGotten != nil && userGotten.IsAdmin {
				proxyReq.Header.Set("x-user-role", "ADMIN")
			}
		}
		httpClient := http.Client{}
		resp, err := httpClient.Do(proxyReq)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		copyHeader(rw.Header(), resp.Header)
		rw.WriteHeader(resp.StatusCode)
		_, _ = rw.Write(response)
	})
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
