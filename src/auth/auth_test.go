package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthMiddleware(t *testing.T) {
	assert := assert.New(t)

	os.Setenv("SESSION_KEY", "not-a-secret-key")
	am := GetAuthMiddleware()

	assert.NotNil(am)
}

func TestAuthMiddleware_LoginHandler(t *testing.T) {
	assert := assert.New(t)

	apiReq := ApiLogin{
		Username: "admin",
		Password: "admin1234",
	}

	data, err := json.Marshal(apiReq)
	if assert.NoError(err) {
		w := httptest.NewRecorder()

		os.Setenv("SESSION_KEY", "not-a-secret-key")
		am := GetAuthMiddleware()

		r := gin.Default()
		r.POST("/user/login", am.LoginHandler)

		req, _ := http.NewRequest("POST", "/user/login", bytes.NewReader(data))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Content-Length", string(len(data)))
		r.ServeHTTP(w, req)

		assert.Equal(http.StatusOK, w.Code)
	}
}

func TestAuthMiddleware_LoginHandler2(t *testing.T) {
	assert := assert.New(t)

	apiReq := ApiLogin{
		Username: "admin",
		Password: "admin12345",
	}

	data, err := json.Marshal(apiReq)
	if assert.NoError(err) {
		w := httptest.NewRecorder()

		os.Setenv("SESSION_KEY", "not-a-secret-key")
		am := GetAuthMiddleware()

		r := gin.Default()
		r.POST("/user/login", am.LoginHandler)

		req, _ := http.NewRequest("POST", "/user/login", bytes.NewReader(data))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Content-Length", string(len(data)))
		r.ServeHTTP(w, req)

		assert.Equal(http.StatusUnauthorized, w.Code)
	}
}