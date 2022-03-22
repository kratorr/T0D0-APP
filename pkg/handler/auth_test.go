package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"todo/models"
	"todo/pkg/service"
	"todo/pkg/service/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// 	"github.com/stretchr/testify/mock"
func TestSignUpHandler(t *testing.T) {
	t.Run("Empty login or password", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		router := gin.Default()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rr := httptest.NewRecorder()
		fmt.Println(rr)

		mockAuthService := mock.NewMockAuth(ctrl)

		userInput := models.User{
			Login:    "test",
			Password: "test",
		}

		mockAuthService.EXPECT().SignUp(userInput).Return(nil)

		// mockAuthService.EXPECT().SignUp(u).Return(nil)
		// SignUp(models.User) error
		// mockAuthService.On("SignUp", mock.AnythingOfType("User")).Return(nil)

		services := &service.Service{mockAuthService}
		handlers := NewHandler(services)

		handlers.InitRoutes(router)

		reqBody, err := json.Marshal(gin.H{
			"login":    "test",
			"password": "test",
		})
		assert.NoError(t, err)
		request, err := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
	})

	// t.Run("Already created user 400 bad request", func(t *testing.T) {})
}
