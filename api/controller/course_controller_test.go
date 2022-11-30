package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mockdb "sugam-project/api/mocks"
	"sugam-project/api/models"
	"sugam-project/api/repository"
)

type CTestServer struct {
	MCourse repository.ICourse
	Router  *mux.Router
}

func newCourseTestServer(t *testing.T, store *CTestServer) *Server {
	server, err := NewCServer(store)
	require.NoError(t, err)
	return server
}

func NewCServer(store *CTestServer) (*Server, error) {
	server := &Server{
		MRepo:  store,
		Router: mux.NewRouter(),
	}
	server.tinitializeRoutes()
	return server, nil
}
func (server *Server) tinitializeRoutes() {
	if value, ok := server.MRepo.(*CTestServer); ok {
		repo = value.MCourse
	}
	server.setJSON("/course", server.CreateCourse, "POST")
	server.setJSON("/course/{id}", server.GetCourseByID, "GET")
	server.setJSON("/courses", server.GetCourses, "GET")
	server.setJSON("/course/{id}", server.UpdateCourse, "PUT")
	server.setJSON("/course/{id}", server.DeleteCourse, "DELETE")

}

func TestGETCourseAPI(t *testing.T) {
	course := &models.Course{
		Name: "domainhem",
	}
	course.ID = 1
	testCases := []struct {
		name          string
		PID           string
		buildStubs    func(cstore *mockdb.MockICourse)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			PID:  "1",
			buildStubs: func(cstore *mockdb.MockICourse) {
				cstore.EXPECT().FindbyId(gomock.Any(), uint(1)).Times(1).Return(course, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BAD_REQUEST",
			PID:  "A",
			buildStubs: func(cstore *mockdb.MockICourse) {
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "NOT_FOUND",
			PID:  "1",
			buildStubs: func(cstore *mockdb.MockICourse) {
				cstore.EXPECT().FindbyId(gomock.Any(), uint(1)).Times(1).Return(&models.Course{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cstore := mockdb.NewMockICourse(ctrl)
			tc.buildStubs(cstore)
			tdnsserver := &CTestServer{
				MCourse: cstore,
			}
			server := newCourseTestServer(t, tdnsserver)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/course/%v", tc.PID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			assert.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(recorder)
		})
	}
}

func TestGETAllCourseAPI(t *testing.T) {
	course := &[]models.Course{
		{
			Name: "domainhem",
		},
		{
			Name: "domainhem",
		},
	}
	testCases := []struct {
		name          string
		buildStubs    func(cstore *mockdb.MockICourse)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildStubs: func(cstore *mockdb.MockICourse) {
				cstore.EXPECT().FindAllCourse(gomock.Any()).Times(1).Return(course, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "NOT_FOUND",
			buildStubs: func(cstore *mockdb.MockICourse) {
				cstore.EXPECT().FindAllCourse(gomock.Any()).Times(1).Return(&[]models.Course{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cstore := mockdb.NewMockICourse(ctrl)
			tc.buildStubs(cstore)
			tdnsserver := &CTestServer{
				MCourse: cstore,
			}
			server := newCourseTestServer(t, tdnsserver)
			recorder := httptest.NewRecorder()

			url := "/courses"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			assert.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(recorder)
		})
	}
}
