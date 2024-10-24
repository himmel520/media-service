package httpctrl

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/mediaAd/internal/usecase"
	"github.com/himmel520/uoffer/mediaAd/internal/usecase/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
			wantRespBody:   `{"message":"invalid id"}`,
		},
		{
			name:           "Err id is not a number",
			id:             "adc",
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"invalid id"}`,
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

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/test-validation/%v", tc.id), nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatusCode, w.Code)
			if tc.wantRespBody != "" {
				assert.Equal(t, tc.wantRespBody, w.Body.String())
			}
		})
	}
}

func TestJwtAdminAccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type args struct {
		token string
		user  string
	}

	type mockBehaviour func(m *mocks.AuthUC, args args)

	testCases := []struct {
		name           string
		args           args
		mockBehaviour  mockBehaviour
		wantStatusCode int
		wantRespBody   string
	}{
		{
			name: "OK",
			args: args{
				token: "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.DkCyanZAGNrDI91Isw6y6VKhZ49ut8vz4Xepbx1WBiFbstSFfxqXPCTXo1tNNYMtoLLStMdx8wQwXzLYuoAVYt0R8O2X6d5Zp7si019vqS-aG_MGD-WX0MPetoQSe8wyA0FHCv487GjZ2uvYwo4mcJNZ-AiuEag6IdlfIQZQlrx7-gUy6pkpZM53K_ynxU1iY55rWAYIbPPZSEXr_JrHMSLU_L5ucNPIpqoWL_w12-w8uxHJ_ithE9LxwCkvgD0Umhoy7FSlg-0Ql_0LXgN1UOi-3o2zq_pxgTELsAAdB3PIhNATfenGGO70_yXk3j_YeeqqrLlaGv_MHd-WM-PSZA",
				user:  "admin",
			},
			mockBehaviour: func(m *mocks.AuthUC, args args) {
				token := strings.TrimPrefix(args.token, "Bearer ")
				m.On("GetUserRoleFromToken", token).Return(args.user, nil).Once()
				m.On("IsUserAdmin", args.user).Return(true).Once()
			},
			wantStatusCode: http.StatusOK,
			wantRespBody:   "",
		},
		{
			name: "invalid user role",
			args: args{
				token: "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciJ9.a0mZwlGQ2mj8gH10l1bDasyh-WE_jOISXV6uYsrL26XuUk1fq3mlMvsoP94R9pgDczP0ksJ7yfBZU4g6Tfjh0EhT1ra47s7ECru8oP_PUzQkMw93RX5WQiw4qai4NhBWmSYStuAaHbPzdBsBqIjE2zp6jgvGZWpsvcsG-Vcyh6z2CaCiViTrvaVx_9mtKAQoZ6rLDy8JULD-wWQdWcv5Sq3Z9IjIb9XI6i52EEZsYMpcSxqvg8KD6iPzWgUqWPKGUJF3BKhy0e-aIWFMa8uIrUBIPm4Nk_hU7SzHUuZztXnzKX_XGttwJVAr-pwqoKJjvfJcYCquV-ybBoozlTxSvQ",
				user:  "invalid",
			},
			mockBehaviour: func(m *mocks.AuthUC, args args) {
				token := strings.TrimPrefix(args.token, "Bearer ")
				m.On("GetUserRoleFromToken", token).Return(args.user, nil).Once()
				m.On("IsUserAdmin", args.user).Return(false).Once()
			},
			wantStatusCode: http.StatusForbidden,
			wantRespBody:   `{"message":"You don't have access to this resource"}`,
		},
		{
			name: "Empty auth header",
			args: args{
				token: "",
			},
			mockBehaviour:  func(m *mocks.AuthUC, args args) {},
			wantStatusCode: http.StatusUnauthorized,
			wantRespBody:   `{"message":"Authorization header is missing"}`,
		},
		{
			name: "invalid auth header format",
			args: args{
				token: "invalid token",
			},
			mockBehaviour:  func(m *mocks.AuthUC, args args) {},
			wantStatusCode: http.StatusUnauthorized,
			wantRespBody:   `{"message":"Authorization header is invalid"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockAuthUsecase := new(mocks.AuthUC)
			h := &Handler{&usecase.Usecase{Auth: mockAuthUsecase}, logrus.New()}

			tc.mockBehaviour(mockAuthUsecase, tc.args)

			router := gin.New()
			router.Use(h.jwtAdminAccess())
			router.GET("/test-auth", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			req := httptest.NewRequest(http.MethodGet, "/test-auth", nil)
			req.Header.Set("Authorization", tc.args.token)

			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatusCode, w.Code)
			if tc.wantRespBody != "" {
				assert.Equal(t, tc.wantRespBody, w.Body.String())
			}
			mockAuthUsecase.AssertExpectations(t)
		})
	}
}

func TestDeleteCategoriesCache(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type mockBehaviour func(m *mocks.AdvUC)

	testCases := []struct {
		name           string
		method         string
		mockBehaviour  mockBehaviour
		wantStatusCode int
	}{
		{
			name:   "OK",
			method: http.MethodPost,
			mockBehaviour: func(m *mocks.AdvUC) {
				m.On("DeleteCache", mock.Anything).Return(nil).Once()
				// time.Sleep(1 * time.Second)
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name:   "GET req should not delete cache",
			method: http.MethodGet,
			mockBehaviour: func(m *mocks.AdvUC) {
				m.On("DeleteCache", mock.Anything).Return(nil).Maybe()
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name:   "Aborted req should not delete cache",
			method: http.MethodPut,
			mockBehaviour: func(m *mocks.AdvUC) {
				m.On("DeleteCache", mock.Anything).Return(nil).Maybe()
			},
			wantStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockAdvUsecase := new(mocks.AdvUC)
			var wg sync.WaitGroup
			h := &Handler{&usecase.Usecase{Adv: mockAdvUsecase}, logrus.New()}

			tc.mockBehaviour(mockAdvUsecase)

			router := gin.New()
			router.Use(h.deleteCategoriesCache(&wg))
			router.GET("/test", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})
			router.POST("/test", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})
			router.PUT("/test", func(c *gin.Context) {
				c.AbortWithStatus(http.StatusBadRequest)
			})

			req := httptest.NewRequest(tc.method, "/test", nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			wg.Wait()

			assert.Equal(t, tc.wantStatusCode, w.Code)
			mockAdvUsecase.AssertExpectations(t)
		})
	}
}
