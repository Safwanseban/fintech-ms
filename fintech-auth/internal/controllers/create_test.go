package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fintechGo/internal/types"
	"fintechGo/internal/usecases/interfaces"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	usecase, _ := Initialize()
	tests := []struct {
		name             string
		expectedStruct   *types.AuthUser
		expectedHttpCode int
		expectedError    error
		assertFunction   func(*httptest.ResponseRecorder)
	}{

		{
			name:             "error- bad payload data ",
			expectedStruct:   nil,
			expectedHttpCode: http.StatusBadRequest,
			expectedError:    errors.New("bad payload data"),
			assertFunction: func(rr *httptest.ResponseRecorder) {

			},
		},
		{
			name: "error- server ",
			expectedStruct: &types.AuthUser{
				Name:     "safwan",
				Email:    "something@gmail.com",
				Password: "12345",
			},
			expectedHttpCode: http.StatusInternalServerError,
			expectedError:    errors.New("internal server error"),
			assertFunction: func(rr *httptest.ResponseRecorder) {

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			data := []byte{}
			var err error
			if tt.expectedStruct != nil {
				data, err = json.Marshal(tt.expectedStruct)
				require.NoError(t, err)

			}
			w := testSetup(http.MethodPost, "/register", bytes.NewBuffer(data), nil, usecase)
			fmt.Println(w)
		},
		)

	}

}
func testSetup(method string, url string,
	body *bytes.Buffer, context map[string]interface{},
	usecase interfaces.UserInterface,

) *httptest.ResponseRecorder {

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, body)
	router := gin.Default()

	NewHandler(router, usecase)
	router.ServeHTTP(w, req)
	return w

}
