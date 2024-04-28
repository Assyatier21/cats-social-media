package helper

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/backendmagang/project-1/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestWriteResponse(t *testing.T) {
	testCases := []struct {
		Name           string
		StandardReq    models.StandardResponseReq
		ExpectedStatus int
		ExpectedBody   string
	}{
		{
			Name: "Success Response",
			StandardReq: models.StandardResponseReq{
				Code:    http.StatusOK,
				Message: "OK",
				Data:    "example data",
			},
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"code":200,"status":"success","message":"OK","data":"example data","error":null}`,
		},
		{
			Name: "Failed Response",
			StandardReq: models.StandardResponseReq{
				Code:    http.StatusBadRequest,
				Message: "Bad Request",
				Error:   errors.New("invalid request"),
			},
			ExpectedStatus: http.StatusBadRequest,
			ExpectedBody:   `{"code":400,"status":"failed","message":"Bad Request","data":null,"error":"invalid request"}`,
		},
		{
			Name: "Empty Message",
			StandardReq: models.StandardResponseReq{
				Code:    http.StatusBadRequest,
				Message: "",
				Error:   errors.New("invalid request"),
			},
			ExpectedStatus: http.StatusBadRequest,
			ExpectedBody:   `{"code":400,"status":"failed","message":"Bad Request","data":null,"error":"invalid request"}`,
		},
		// Add more test cases here...
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Create a new Echo context for testing
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Call the WriteResponse function
			err := WriteResponse(c, tc.StandardReq)
			assert.NoError(t, err)

			// Check the response status code
			assert.Equal(t, tc.ExpectedStatus, rec.Code)

			// Remove the newline character from the expected body
			expectedBody := strings.TrimSpace(tc.ExpectedBody)

			// Compare the response body after trimming whitespace
			assert.Equal(t, expectedBody, strings.TrimSpace(rec.Body.String()))

			// Compare the Content-Type header while ignoring letter casing
			assert.Equal(t, "application/json; charset=utf-8", strings.ToLower(rec.Header().Get("Content-Type")))
		})
	}
}
