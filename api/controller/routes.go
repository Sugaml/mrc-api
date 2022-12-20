package controller

import (
	"net/http"
	"sugam-project/api/middleware"
	"sugam-project/api/responses"
)

func (server *Server) setJSON(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.HandleFunc(path, middleware.SetMiddlewareJSON(next)).Methods(method, "OPTIONS")
}

func (server *Server) setAdmin(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.setJSON(path, middleware.SetAdminMiddlewareAuthentication(next), method)
}

func (server *Server) initializeRoutes() {
	server.Router.Use(middleware.CORS)
	server.setJSON("/", server.WelcomePage, "GET")

	server.setJSON("/course", server.CreateCourse, "POST")
	server.setJSON("/course/{id}", server.GetCourseByID, "GET")
	server.setJSON("/courses", server.GetCourses, "GET")
	server.setAdmin("/course/{id}", server.UpdateCourse, "PUT")
	server.setAdmin("/course/{id}", server.DeleteCourse, "DELETE")

	server.setJSON("/user", server.CreateUser, "POST")
	server.setJSON("/user/login", server.GetLogin, "POST")
	server.setJSON("/user", server.GetUserByID, "GET")
	server.setJSON("/users", server.GetUsers, "GET")
	server.setJSON("/user/{id}", server.UpdateUser, "PUT")
	server.setJSON("/user/{id}", server.DeleteUser, "DELETE")
	server.setJSON("/user/forgot-password", server.ForgotPassword, "POST")

	server.setJSON("/student_info", server.StudentInfo, "POST")
	server.setJSON("/student_info/{id}", server.StudentDetail, "GET")
	server.setJSON("/student/general-info", server.StudentGeneralInfo, "GET")
	server.setJSON("/student/address", server.StudentAddress, "POST")
	server.setJSON("/student/education", server.StudentEducation, "POST")
	server.setJSON("/student/file", server.StudentFileInfo, "POST")
	server.setJSON("/student/{sid}/address", server.GetStudentAddress, "GET")
	server.setJSON("/student/{sid}/document", server.GetStudentFile, "GET")

	server.setJSON("/student/{sid}", server.UpdateStudentInfo, "PUT")
	server.setJSON("/student/{sid}/address", server.UpdateStudentAddress, "PUT")
	server.setJSON("/student/{sid}/education", server.UpdateStudentEducation, "PUT")
	server.setJSON("/student/{sid}/document", server.UpdateStudentFile, "PUT")
}

func (server *Server) WelcomePage(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "welcome to mrc-api service")
}
