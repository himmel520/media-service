package httpctrl

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/mediaAd/internal/controller"
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
			wantRespBody:   fmt.Sprintf(`{"message":"%v"}`, controller.ErrInvalidID),
		},
		{
			name:           "Err id is negative",
			id:             "-1",
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   fmt.Sprintf(`{"message":"%v"}`, controller.ErrInvalidID),
		},
		{
			name:           "Err id is not a number",
			id:             "adc",
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   fmt.Sprintf(`{"message":"%v"}`, controller.ErrInvalidID),
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
		// {
			// 	name: "OK",
			// 	args: args{
			// 		token: "Bearer TODO: Брать безопасно",
			// 		user:  "admin",
			// 	},
			// 	mockBehaviour: func(m *mocks.AuthUC, args args) {
			// 		token := strings.TrimPrefix(args.token, "Bearer ")
			// 		m.On("GetUserRoleFromToken", token).Return(args.user, nil).Once()
			// 		m.On("IsUserAdmin", args.user).Return(true).Once()
			// 	},
			// 	wantStatusCode: http.StatusOK,
			// 	wantRespBody:   "",
			// },
			// {
			// 	name: "invalid user role",
			// 	args: args{
			// 		token: "Bearer TODO: Брать безопасно",
			// 		user:  "invalid",
			// 	},
			// 	mockBehaviour: func(m *mocks.AuthUC, args args) {
			// 		token := strings.TrimPrefix(args.token, "Bearer ")
			// 		m.On("GetUserRoleFromToken", token).Return(args.user, nil).Once()
			// 		m.On("IsUserAdmin", args.user).Return(false).Once()
			// 	},
			// 	wantStatusCode: http.StatusForbidden,
			// 	wantRespBody:   fmt.Sprintf(`{"message":"%v"}`, controller.ErrForbidden),
		// },
		{
			name: "Empty auth header",
			args: args{
				token: "",
			},
			mockBehaviour:  func(m *mocks.AuthUC, args args) {},
			wantStatusCode: http.StatusUnauthorized,
			wantRespBody:   fmt.Sprintf(`{"message":"%v"}`, controller.ErrEmptyAuthHeader),
		},
		{
			name: "invalid auth header format",
			args: args{
				token: "invalid token",
			},
			mockBehaviour:  func(m *mocks.AuthUC, args args) {},
			wantStatusCode: http.StatusUnauthorized,
			wantRespBody:   fmt.Sprintf(`{"message":"%v"}`, controller.ErrInvalidAuthHeader),
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
			h := New(&usecase.Usecase{Adv: mockAdvUsecase}, logrus.New())

			tc.mockBehaviour(mockAdvUsecase)

			router := gin.New()
			router.Use(h.deleteCategoriesCache())
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

			assert.Equal(t, tc.wantStatusCode, w.Code)
			mockAdvUsecase.AssertExpectations(t)
		})
	}
}
