package httpctrl

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidateID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []struct {
		name           string
		id             string
		wantStatusCode int
		wantRespBody   string
	}{
		{
			name:           "Ok",
			id:             "1",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "Err id is 0",
			id:             "0",
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"invalid id"}`,
		},
		{
			name:           "Err id is negative",
			id:             "-1",
			wantStatusCode: http.StatusBadRequest,
			wantRespBody: `{"message":"invalid id"}`,
		},
		{
			name:           "Err id is not a number",
			id:             "adc",
			wantStatusCode: http.StatusBadRequest,
			wantRespBody: `{"message":"invalid id"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h := &Handler{}

			router := gin.Default()
			router.Use(h.validateID())
			router.GET("/test-validation/:id", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/test-validation/%v", tc.id), nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatusCode, w.Code)
			if tc.wantRespBody != "" {
				assert.Equal(t, tc.wantRespBody, w.Body.String())
			}
		})
	}

}
