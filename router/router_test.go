package router_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"web-service-gin/auth/signup"
	"web-service-gin/model"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

// func TestAlbumController(t *testing.T) {
// 	type album struct {
// 		Id     string
// 		Title  string
// 		Artist string
// 		Price  float64
// 	}

// 	var albums = []album{
// 		{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 		{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 		{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// 	}
// 	r := router.InitRoute()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodGet, "/api/album", nil)
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, albums, w.Body)
// }

// type Tests struct {
// 	name          string
// 	server        *httptest.Server
// 	response      []byte
// 	expectedError error
// }

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestSignup(t *testing.T) {
	mockResponse := `{"message":"User Created"}`
	r := SetupRouter()
	mockRequestBody := model.UserDetails{
		Username: "testuser",
		Password: "testpassword",
	}
	jsonValue, _ := json.Marshal(mockRequestBody)
	r.POST("/signup", signup.Signup)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, responseData, mockResponse)
	assert.Equal(t, http.StatusCreated, w.Code)
	//	tests := []Tests{
	//		{
	//			name: "test basic request",
	//			server: http.NewServer(gin.HandlerFunc(gin.New().FuncMap))),
	//		}
	//	}
}
