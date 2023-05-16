package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"todo/models"
	"todo/pkg/service/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// 	"github.com/stretchr/testify/mock"

type args struct{}

func TestSignUpUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name           string
		args           args
		beforeTest     func(SignUp *mock.MockAuth)
		body           map[string]interface{}
		wantStatusCode int
		wantHeader     http.Header
		wantBody       string
	}{
		{
			name: "Create valid user",
			beforeTest: func(SignUp *mock.MockAuth) {
				SignUp.EXPECT().SignUp(
					models.User{
						Password: "test_password",
						Login:    "test_login",
					},
				).Return(nil)
			},
			body:           map[string]interface{}{"login": "test_login", "password": "test_password"},
			wantStatusCode: 200,
			wantHeader:     http.Header{"Content-Type": {"application/json; charset=utf-8"}},
			wantBody:       `{"response":"user created"}`,
		},
	}

	for _, tcase := range tests {
		t.Run(tcase.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.Default()
			mockAuthService := mock.NewMockAuth(ctrl)

			rr := httptest.NewRecorder()

			//	services := &service.Service{mockAuthService}
			//	handlers := NewHandler(services)

			//	handlers.InitRoutes(router)

			if tcase.beforeTest != nil {
				tcase.beforeTest(mockAuthService)
			}
			reqBody, err := json.Marshal(tcase.body)

			assert.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)

			router.ServeHTTP(rr, request)

			if !reflect.DeepEqual(rr.Result().StatusCode, tcase.wantStatusCode) {
				t.Errorf("SignUp() = %v, want %v", rr.Result().StatusCode, tcase.wantStatusCode)
			}

			// if !reflect.DeepEqual(rr.Header, tcase.wantHeader) {
			// 	t.Errorf("SignUp() = %v, want %v", rr.Header(), tcase.wantHeader)
			// }

			bodyBuffer := new(bytes.Buffer)
			bodyBuffer.ReadFrom(rr.Body)
			body := strings.TrimSpace(bodyBuffer.String())

			if !reflect.DeepEqual(body, tcase.wantBody) {
				t.Errorf("SignUp() = %s, want %s", bodyBuffer.String(), tcase.wantBody)
			}
		})
	}
}
