package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/pkg/service"
	"todo/pkg/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignUpHandler(t *testing.T) {
	t.Run("Register valid user", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		rr := httptest.NewRecorder()
		mockAuthService := new(mocks.Auth)
		mockAuthService.On("SignUp", mock.AnythingOfType("User")).Return(nil)
		services := &service.Service{mockAuthService}
		handlers := NewHandler(services)

		handlers.InitRoutes(router)

		reqBody, err := json.Marshal(gin.H{
			"login":    "VS_Code",
			"password": "q",
		})
		assert.NoError(t, err)
		request, err := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 200, rr.Code)
		mockAuthService.AssertNotCalled(t, "SignUp")
	})

	// t.Run("Already created user 400 bad request", func(t *testing.T) {})
}
