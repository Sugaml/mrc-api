package controller

import (
	"net/http"
	"sugam-project/api/middleware"
	"sugam-project/api/responses"
)

func (server *Server) setJSON(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.HandleFunc(path, middleware.SetMiddlewareJSON(next)).Methods(method, "OPTIONS")
}

// func (server *Server) setAdmin(path string, next func(http.ResponseWriter, *http.Request), method string) {
// 	server.setJSON(path, middleware.SetAdminMiddlewareAuthentication(next), method)
// }

func (server *Server) initializeRoutes() {
	server.Router.Use(middleware.CORS)
	server.setJSON("/", server.WelcomePage, "GET")

	server.setJSON("/course", server.CreateCourse, "POST")
	server.setJSON("/course/{id}", server.GetCourseByID, "GET")
	server.setJSON("/courses", server.GetCourses, "GET")
	server.setJSON("/course/{id}", server.UpdateCourse, "PUT")
	server.setJSON("/course/{id}", server.DeleteCourse, "DELETE")

	server.setJSON("/user", server.CreateUser, "POST")
	server.setJSON("/user/login", server.GetLogin, "POST")
	server.setJSON("/user/{id}", server.GetUserByID, "GET")
	server.setJSON("/users", server.GetUsers, "GET")
	server.setJSON("/user/{id}", server.UpdateUser, "PUT")
	server.setJSON("/user/{id}", server.DeleteUser, "DELETE")

	server.setJSON("/user/forgot-password", server.ForgotPassword, "POST")

	server.setJSON("/student_info", server.StudentInfo, "POST")
}

func (server *Server) WelcomePage(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "welcome to sms-api project")
}
