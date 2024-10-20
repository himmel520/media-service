package httpctrl

// import (
// 	"bytes"
// 	"context"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/himmel520/uoffer/mediaAd/internal/models"
// 	"github.com/himmel520/uoffer/mediaAd/internal/repository/repoerr"
// 	"github.com/himmel520/uoffer/mediaAd/internal/usecase/mocks"

// 	"github.com/sirupsen/logrus"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetLogo(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	type args struct {
// 		id int
// 	}

// 	type mockBehaviour func(m *mocks.LogoSrv, args args)

// 	testCases := []struct {
// 		name           string
// 		args           args
// 		mockBehaviour  mockBehaviour
// 		wantStatusCode int
// 		wantRespBody   string
// 	}{
// 		{
// 			name: "OK",
// 			args: args{
// 				id: 1,
// 			},
// 			mockBehaviour: func(m *mocks.LogoSrv, args args) {
// 				expectedLogo := &models.LogoResp{
// 					ID:    1,
// 					Url:   "http://example.com/logo.png",
// 					Title: "Test Logo",
// 				}
// 				m.On("GetByID", context.Background(), args.id).Return(expectedLogo, nil)
// 			},
// 			wantStatusCode: http.StatusOK,
// 			wantRespBody:   `{"id":1,"url":"http://example.com/logo.png","title":"Test Logo"}`,
// 		},
// 		{
// 			name: "Logo Not Found",
// 			args: args{
// 				id: 999,
// 			},
// 			mockBehaviour: func(m *mocks.LogoSrv, args args) {
// 				m.On("GetByID", context.Background(), args.id).Return(nil, repoerr.ErrLogoNotFound)
// 			},
// 			wantStatusCode: http.StatusNotFound,
// 			wantRespBody:   `{"message":"logo not found"}`,
// 		},
// 		{
// 			name: "Internal Server Error",
// 			args: args{
// 				id: 2,
// 			},
// 			mockBehaviour: func(m *mocks.LogoSrv, args args) {
// 				m.On("GetByID", context.Background(), 2).Return(nil, errors.New("internal server error"))
// 			},
// 			wantStatusCode: http.StatusInternalServerError,
// 			wantRespBody:   `{"message":"internal server error"}`,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			mockLogoService := new(mocks.LogoSrv)
// 			handler := &Handler{&service.Service{Logo: mockLogoService}, logrus.New()}

// 			tc.mockBehaviour(mockLogoService, tc.args)

// 			router := gin.Default()
// 			router.GET("/logos/:id", handler.getLogo)

// 			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/logos/%v", tc.args.id), nil)
// 			w := httptest.NewRecorder()

// 			router.ServeHTTP(w, req)

// 			assert.Equal(t, tc.wantStatusCode, w.Code)
// 			assert.JSONEq(t, tc.wantRespBody, w.Body.String())
// 			mockLogoService.AssertExpectations(t)
// 		})
// 	}
// }

// func TestGetLogos(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	type mockBehaviour func(m *mocks.LogoSrv)

// 	testCases := []struct {
// 		name           string
// 		mockBehaviour  mockBehaviour
// 		wantStatusCode int
// 		wantRespBody   string
// 	}{
// 		{
// 			name: "OK",
// 			mockBehaviour: func(m *mocks.LogoSrv) {
// 				expectedLogos := []*models.LogoResp{
// 					{
// 						ID:    1,
// 						Url:   "http://example.com/logo.png",
// 						Title: "Test Logo",
// 					},
// 					{
// 						ID:    2,
// 						Url:   "http://example.com/logo.png",
// 						Title: "Test Logo",
// 					},
// 				}

// 				m.On("GetAll", context.Background()).Return(expectedLogos, nil)
// 			},
// 			wantStatusCode: http.StatusOK,
// 			wantRespBody:   `[{"id":1,"url":"http://example.com/logo.png","title":"Test Logo"},{"id":2,"url":"http://example.com/logo.png","title":"Test Logo"}]`,
// 		},
// 		{
// 			name: "Logo Not Found",
// 			mockBehaviour: func(m *mocks.LogoSrv) {
// 				m.On("GetAll", context.Background()).Return(nil, repoerr.ErrLogoNotFound)
// 			},
// 			wantStatusCode: http.StatusNotFound,
// 			wantRespBody:   `{"message":"logo not found"}`,
// 		},
// 		{
// 			name: "Internal Server Error",
// 			mockBehaviour: func(m *mocks.LogoSrv) {
// 				m.On("GetAll", context.Background()).Return(nil, errors.New("internal server error"))
// 			},
// 			wantStatusCode: http.StatusInternalServerError,
// 			wantRespBody:   `{"message":"internal server error"}`,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			mockLogoService := new(mocks.LogoSrv)
// 			handler := &Handler{&service.Service{Logo: mockLogoService}, logrus.New()}

// 			tc.mockBehaviour(mockLogoService)

// 			router := gin.Default()
// 			router.GET("/logos", handler.getLogos)

// 			req, _ := http.NewRequest(http.MethodGet, "/logos", nil)
// 			w := httptest.NewRecorder()

// 			router.ServeHTTP(w, req)

// 			assert.Equal(t, tc.wantStatusCode, w.Code)
// 			assert.JSONEq(t, tc.wantRespBody, w.Body.String())
// 			mockLogoService.AssertExpectations(t)

// 		})
// 	}
// }

// func TestAddLogo(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	type args struct {
// 		in *models.Logo
// 	}

// 	type mockBehaviour func(m *mocks.LogoSrv, args args)

// 	testCases := []struct {
// 		name           string
// 		args           args
// 		mockBehaviour  mockBehaviour
// 		inputbody      string
// 		wantStatusCode int
// 		wantRespBody   string
// 	}{
// 		{
// 			name: "OK",
// 			args: args{
// 				in: &models.Logo{
// 					Url:   "http://example.com/logo.png",
// 					Title: "Test Logo",
// 				},
// 			},
// 			mockBehaviour: func(m *mocks.LogoSrv, args args) {
// 				expectedLogo := &models.LogoResp{
// 					ID:    1,
// 					Url:   "http://example.com/logo.png",
// 					Title: "Test Logo",
// 				}
// 				m.On("Add", context.Background(), args.in).Return(expectedLogo, nil)
// 			},
// 			inputbody:      `{"url":"http://example.com/logo.png","title":"Test Logo"}`,
// 			wantStatusCode: http.StatusCreated,
// 			wantRespBody:   `{"id":1,"url":"http://example.com/logo.png","title":"Test Logo"}`,
// 		},
// 		// Repository error
// 		{
// 			name: "Err logo exist",
// 			args: args{
// 				in: &models.Logo{
// 					Url:   "http://example.com/logo.png",
// 					Title: "Test Logo",
// 				},
// 			},
// 			mockBehaviour: func(m *mocks.LogoSrv, args args) {
// 				m.On("Add", context.Background(), args.in).Return(nil, repoerr.ErrLogoExist)
// 			},
// 			inputbody:      `{"url":"http://example.com/logo.png","title":"Test Logo"}`,
// 			wantStatusCode: http.StatusBadRequest,
// 			wantRespBody:   `{"message":"logo url must be unique"}`,
// 		},
// 		{
// 			name: "Internal Server Error",
// 			args: args{
// 				in: &models.Logo{
// 					Url:   "http://example.com/logo.png",
// 					Title: "Test Logo",
// 				},
// 			},
// 			mockBehaviour: func(m *mocks.LogoSrv, args args) {
// 				m.On("Add", context.Background(), args.in).Return(nil, errors.New("internal server error"))
// 			},
// 			inputbody:      `{"url":"http://example.com/logo.png","title":"Test Logo"}`,
// 			wantStatusCode: http.StatusInternalServerError,
// 			wantRespBody:   `{"message":"internal server error"}`,
// 		},
// 		// Validation
// 		{
// 			name:           "Invalid resp body",
// 			inputbody:      `{}`,
// 			args:           args{},
// 			mockBehaviour:  func(m *mocks.LogoSrv, args args) {},
// 			wantStatusCode: http.StatusBadRequest,
// 			wantRespBody:   `{"message":"Key: 'Logo.Url' Error:Field validation for 'Url' failed on the 'required' tag\nKey: 'Logo.Title' Error:Field validation for 'Title' failed on the 'required' tag"}`,
// 		},
// 		{
// 			name:           "Invalid url: length is less than 3",
// 			inputbody:      `{"url":"ht","title":"Test Logo"}`,
// 			args:           args{},
// 			mockBehaviour:  func(m *mocks.LogoSrv, args args) {},
// 			wantStatusCode: http.StatusBadRequest,
// 			wantRespBody:   `{"message":"Key: 'Logo.Url' Error:Field validation for 'Url' failed on the 'min' tag"}`,
// 		},
// 		{
// 			name:           "Invalid url: missing url",
// 			inputbody:      `{"title":"Test Logo"}`,
// 			args:           args{},
// 			mockBehaviour:  func(m *mocks.LogoSrv, args args) {},
// 			wantStatusCode: http.StatusBadRequest,
// 			wantRespBody:   `{"message":"Key: 'Logo.Url' Error:Field validation for 'Url' failed on the 'required' tag"}`,
// 		},
// 		{
// 			name:           "Invalid title: length is less than 3",
// 			inputbody:      `{"url":"http://example.com/logo.png","title":"Te"}`,
// 			args:           args{},
// 			mockBehaviour:  func(m *mocks.LogoSrv, args args) {},
// 			wantStatusCode: http.StatusBadRequest,
// 			wantRespBody:   `{"message":"Key: 'Logo.Title' Error:Field validation for 'Title' failed on the 'min' tag"}`,
// 		},
// 		{
// 			name:           "Invalid title: missing title",
// 			inputbody:      `{"url":"http://example.com/logo.png"}`,
// 			args:           args{},
// 			mockBehaviour:  func(m *mocks.LogoSrv, args args) {},
// 			wantStatusCode: http.StatusBadRequest,
// 			wantRespBody:   `{"message":"Key: 'Logo.Title' Error:Field validation for 'Title' failed on the 'required' tag"}`,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			mockLogoService := new(mocks.LogoSrv)
// 			handler := &Handler{&service.Service{Logo: mockLogoService}, logrus.New()}

// 			tc.mockBehaviour(mockLogoService, tc.args)

// 			router := gin.Default()
// 			router.POST("/logos", handler.addLogo)

// 			req, _ := http.NewRequest(http.MethodPost, "/logos", bytes.NewBufferString(tc.inputbody))
// 			w := httptest.NewRecorder()

// 			router.ServeHTTP(w, req)

// 			assert.Equal(t, tc.wantStatusCode, w.Code)
// 			assert.JSONEq(t, tc.wantRespBody, w.Body.String())
// 			mockLogoService.AssertExpectations(t)
// 		})
// 	}
// }

// // func TestUpdateLogo(t *testing.T){
// // 	gin.SetMode(gin.TestMode)

// // 	type args struct {
// // 		id int
// // 		in *models.LogoUpdate
// // 	}

// // 	type mockBehaviour func(m *mocks.LogoSrv, args args)

// // }
