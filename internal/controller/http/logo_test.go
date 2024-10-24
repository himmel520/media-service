package httpctrl

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/uoffer/mediaAd/internal/usecase"
	"github.com/himmel520/uoffer/mediaAd/internal/usecase/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetLogo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type args struct {
		id int
	}

	type mockBehaviour func(m *mocks.LogoUC, args args)

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
				id: 1,
			},
			mockBehaviour: func(m *mocks.LogoUC, args args) {
				expectedLogo := &entity.LogoResp{
					ID:    1,
					Url:   "http://example.com/logo.png",
					Title: "Test Logo",
				}
				m.On("GetByID", context.Background(), args.id).Return(expectedLogo, nil)
			},
			wantStatusCode: http.StatusOK,
			wantRespBody:   `{"id":1,"url":"http://example.com/logo.png","title":"Test Logo"}`,
		},
		{
			name: "Logo Not Found",
			args: args{
				id: 999,
			},
			mockBehaviour: func(m *mocks.LogoUC, args args) {
				m.On("GetByID", context.Background(), args.id).Return(nil, repoerr.ErrLogoNotFound)
			},
			wantStatusCode: http.StatusNotFound,
			wantRespBody:   `{"message":"logo not found"}`,
		},
		{
			name: "Internal Server Error",
			args: args{
				id: 2,
			},
			mockBehaviour: func(m *mocks.LogoUC, args args) {
				m.On("GetByID", context.Background(), 2).Return(nil, errors.New("internal server error"))
			},
			wantStatusCode: http.StatusInternalServerError,
			wantRespBody:   `{"message":"internal server error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockLogoUsecase := new(mocks.LogoUC)
			handler := &Handler{
				uc:  &usecase.Usecase{Logo: mockLogoUsecase},
				log: logrus.New(),
			}

			tc.mockBehaviour(mockLogoUsecase, tc.args)

			router := gin.Default()
			router.GET("/logos/:id", handler.getLogo)

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/logos/%v", tc.args.id), nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatusCode, w.Code)
			assert.JSONEq(t, tc.wantRespBody, w.Body.String())
			mockLogoUsecase.AssertExpectations(t)
		})
	}
}

func TestGetLogos(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type mockBehaviour func(m *mocks.LogoUC)

	testCases := []struct {
		name           string
		mockBehaviour  mockBehaviour
		wantStatusCode int
		wantRespBody   string
	}{
		{
			name: "OK",
			mockBehaviour: func(m *mocks.LogoUC) {
				expectedLogos := []*entity.LogoResp{
					{
						ID:    1,
						Url:   "http://example.com/logo.png",
						Title: "Test Logo",
					},
					{
						ID:    2,
						Url:   "http://example.com/logo.png",
						Title: "Test Logo",
					},
				}

				m.On("GetAll", context.Background()).Return(expectedLogos, nil)
			},
			wantStatusCode: http.StatusOK,
			wantRespBody:   `[{"id":1,"url":"http://example.com/logo.png","title":"Test Logo"},{"id":2,"url":"http://example.com/logo.png","title":"Test Logo"}]`,
		},
		// Errors
		{
			name: "Logo Not Found",
			mockBehaviour: func(m *mocks.LogoUC) {
				m.On("GetAll", context.Background()).Return(nil, repoerr.ErrLogoNotFound)
			},
			wantStatusCode: http.StatusNotFound,
			wantRespBody:   `{"message":"logo not found"}`,
		},
		{
			name: "Internal Server Error",
			mockBehaviour: func(m *mocks.LogoUC) {
				m.On("GetAll", context.Background()).Return(nil, errors.New("internal server error"))
			},
			wantStatusCode: http.StatusInternalServerError,
			wantRespBody:   `{"message":"internal server error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockLogoUsecase := new(mocks.LogoUC)
			handler := &Handler{
				uc:  &usecase.Usecase{Logo: mockLogoUsecase},
				log: logrus.New(),
			}

			tc.mockBehaviour(mockLogoUsecase)

			router := gin.Default()
			router.GET("/logos", handler.getLogos)

			req, _ := http.NewRequest(http.MethodGet, "/logos", nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatusCode, w.Code)
			assert.JSONEq(t, tc.wantRespBody, w.Body.String())
			mockLogoUsecase.AssertExpectations(t)

		})
	}
}

func TestAddLogo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type args struct {
		in *entity.Logo
	}

	type mockBehaviour func(m *mocks.LogoUC, args args)

	testCases := []struct {
		name           string
		args           args
		mockBehaviour  mockBehaviour
		inputbody      string
		wantStatusCode int
		wantRespBody   string
	}{
		{
			name: "OK",
			args: args{
				in: &entity.Logo{
					Url:   "http://example.com/logo.png",
					Title: "Test Logo",
				},
			},
			mockBehaviour: func(m *mocks.LogoUC, args args) {
				expectedLogo := &entity.LogoResp{
					ID:    1,
					Url:   "http://example.com/logo.png",
					Title: "Test Logo",
				}
				m.On("Add", context.Background(), args.in).Return(expectedLogo, nil)
			},
			inputbody:      `{"url":"http://example.com/logo.png","title":"Test Logo"}`,
			wantStatusCode: http.StatusCreated,
			wantRespBody:   `{"id":1,"url":"http://example.com/logo.png","title":"Test Logo"}`,
		},
		// Errors
		{
			name: "Err logo exist",
			args: args{
				in: &entity.Logo{
					Url:   "http://example.com/logo.png",
					Title: "Test Logo",
				},
			},
			mockBehaviour: func(m *mocks.LogoUC, args args) {
				m.On("Add", context.Background(), args.in).Return(nil, repoerr.ErrLogoExist)
			},
			inputbody:      `{"url":"http://example.com/logo.png","title":"Test Logo"}`,
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"logo url must be unique"}`,
		},
		{
			name: "Internal Server Error",
			args: args{
				in: &entity.Logo{
					Url:   "http://example.com/logo.png",
					Title: "Test Logo",
				},
			},
			mockBehaviour: func(m *mocks.LogoUC, args args) {
				m.On("Add", context.Background(), args.in).Return(nil, errors.New("internal server error"))
			},
			inputbody:      `{"url":"http://example.com/logo.png","title":"Test Logo"}`,
			wantStatusCode: http.StatusInternalServerError,
			wantRespBody:   `{"message":"internal server error"}`,
		},
		// Validation
		{
			name:           "Empty resp body",
			inputbody:      `{}`,
			args:           args{},
			mockBehaviour:  func(m *mocks.LogoUC, args args) {},
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"Key: 'Logo.Url' Error:Field validation for 'Url' failed on the 'required' tag\nKey: 'Logo.Title' Error:Field validation for 'Title' failed on the 'required' tag"}`,
		},
		{
			name:           "Invalid url: length is less than 3",
			inputbody:      `{"url":"ht","title":"Test Logo"}`,
			args:           args{},
			mockBehaviour:  func(m *mocks.LogoUC, args args) {},
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"Key: 'Logo.Url' Error:Field validation for 'Url' failed on the 'min' tag"}`,
		},
		{
			name:           "Invalid url: missing url",
			inputbody:      `{"title":"Test Logo"}`,
			args:           args{},
			mockBehaviour:  func(m *mocks.LogoUC, args args) {},
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"Key: 'Logo.Url' Error:Field validation for 'Url' failed on the 'required' tag"}`,
		},
		{
			name:           "Invalid title: length is less than 3",
			inputbody:      `{"url":"http://example.com/logo.png","title":"Te"}`,
			args:           args{},
			mockBehaviour:  func(m *mocks.LogoUC, args args) {},
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"Key: 'Logo.Title' Error:Field validation for 'Title' failed on the 'min' tag"}`,
		},
		{
			name:           "Invalid title: length is more than 100",
			inputbody:      `{"url":"http://example.com/logo.png","title":"Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit adipisci velit..."}`,
			args:           args{},
			mockBehaviour:  func(m *mocks.LogoUC, args args) {},
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"Key: 'Logo.Title' Error:Field validation for 'Title' failed on the 'max' tag"}`,
		},
		{
			name:           "Invalid title: missing title",
			inputbody:      `{"url":"http://example.com/logo.png"}`,
			args:           args{},
			mockBehaviour:  func(m *mocks.LogoUC, args args) {},
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"Key: 'Logo.Title' Error:Field validation for 'Title' failed on the 'required' tag"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockLogoUsecase := new(mocks.LogoUC)
			handler := &Handler{
				uc:  &usecase.Usecase{Logo: mockLogoUsecase},
				log: logrus.New(),
			}

			tc.mockBehaviour(mockLogoUsecase, tc.args)

			router := gin.Default()
			router.POST("/logos", handler.addLogo)

			req, _ := http.NewRequest(http.MethodPost, "/logos", bytes.NewBufferString(tc.inputbody))
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatusCode, w.Code)
			assert.JSONEq(t, tc.wantRespBody, w.Body.String())
			mockLogoUsecase.AssertExpectations(t)
		})
	}
}

func TestUpdateLogo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	adr := func(s string) *string { return &s }

	type args struct {
		id int
		in *entity.LogoUpdate
	}

	type mockBehaviour func(m *mocks.LogoUC, args args)

	testCases := []struct {
		name           string
		args           args
		mockBehaviour  mockBehaviour
		inputbody      string
		wantStatusCode int
		wantRespBody   string
	}{
		{
			name: "OK",
			args: args{
				id: 1,
				in: &entity.LogoUpdate{
					Url:   adr("http://example.com/logo.png"),
					Title: adr("Test Logo"),
				},
			},
			mockBehaviour: func(m *mocks.LogoUC, args args) {
				expectedLogo := &entity.LogoResp{
					ID:    1,
					Url:   "http://example.com/logo.png",
					Title: "Test Logo",
				}
				m.On("Update", context.Background(), args.id, args.in).Return(expectedLogo, nil)
			},
			inputbody:      `{"url":"http://example.com/logo.png","title":"Test Logo"}`,
			wantStatusCode: http.StatusOK,
			wantRespBody:   `{"id":1,"url":"http://example.com/logo.png","title":"Test Logo"}`,
		},
		// Errors
		{
			name: "Err Logo exist",
			args: args{
				id: 1,
				in: &entity.LogoUpdate{
					Url: adr("http://example.com/logo.png"),
				},
			},
			mockBehaviour: func(m *mocks.LogoUC, args args) {
				m.On("Update", context.Background(), args.id, args.in).Return(nil, repoerr.ErrLogoExist)
			},
			inputbody:      `{"url":"http://example.com/logo.png"}`,
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"logo url must be unique"}`,
		},
		{
			name: "Err Logo not found",
			args: args{
				id: 999,
				in: &entity.LogoUpdate{
					Url: adr("http://example.com/logo.png"),
				},
			},
			mockBehaviour: func(m *mocks.LogoUC, args args) {
				m.On("Update", context.Background(), args.id, args.in).Return(nil, repoerr.ErrLogoNotFound)
			},
			inputbody:      `{"url":"http://example.com/logo.png"}`,
			wantStatusCode: http.StatusNotFound,
			wantRespBody:   `{"message":"logo not found"}`,
		},
		{
			name: "Internal server error",
			args: args{
				id: 1,
				in: &entity.LogoUpdate{
					Url: adr("http://example.com/logo.png"),
				},
			},
			mockBehaviour: func(m *mocks.LogoUC, args args) {
				m.On("Update", context.Background(), args.id, args.in).Return(nil, errors.New("Internal server error"))
			},
			inputbody:      `{"url":"http://example.com/logo.png"}`,
			wantStatusCode: http.StatusInternalServerError,
			wantRespBody:   `{"message":"Internal server error"}`,
		},
		// Validation
		{
			name: "Empty resp body",
			args: args{
				id: 1,
			},
			mockBehaviour:  func(m *mocks.LogoUC, args args) {},
			inputbody:      `{}`,
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"logo has no changes"}`,
		},
		{
			name: "Missing url",
			args: args{
				id: 1,
				in: &entity.LogoUpdate{
					Title: adr("Test Logo"),
				},
			},
			mockBehaviour: func(m *mocks.LogoUC, args args) {
				expectedLogo := &entity.LogoResp{
					ID:    1,
					Url:   "http://example.com/logo.png",
					Title: "Test Logo",
				}
				m.On("Update", context.Background(), args.id, args.in).Return(expectedLogo, nil)
			},
			inputbody:      `{"title":"Test Logo"}`,
			wantStatusCode: http.StatusOK,
			wantRespBody:   `{"id":1,"url":"http://example.com/logo.png","title":"Test Logo"}`,
		},
		{
			name:           "Invalid url: length is less than 3",
			args:           args{},
			mockBehaviour:  func(m *mocks.LogoUC, args args) {},
			inputbody:      `{"url":"ht","title":"Test Logo"}`,
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"Key: 'LogoUpdate.Url' Error:Field validation for 'Url' failed on the 'min' tag"}`,
		},
		{
			name: "Missing title",
			args: args{
				id: 1,
				in: &entity.LogoUpdate{
					Url: adr("http://example.com/logo.png"),
				},
			},
			mockBehaviour: func(m *mocks.LogoUC, args args) {
				expectedLogo := &entity.LogoResp{
					ID:    1,
					Url:   "http://example.com/logo.png",
					Title: "Test Logo",
				}
				m.On("Update", context.Background(), args.id, args.in).Return(expectedLogo, nil)
			},
			inputbody:      `{"url":"http://example.com/logo.png"}`,
			wantStatusCode: http.StatusOK,
			wantRespBody:   `{"id":1,"url":"http://example.com/logo.png","title":"Test Logo"}`,
		},
		{
			name:           "Invalid title: length is more than 100",
			args:           args{},
			mockBehaviour:  func(m *mocks.LogoUC, args args) {},
			inputbody:      `{"url":"http://example.com/logo.png","title":"Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit adipisci velit..."}`,
			wantStatusCode: http.StatusBadRequest,
			wantRespBody:   `{"message":"Key: 'LogoUpdate.Title' Error:Field validation for 'Title' failed on the 'max' tag"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockLogoUsecase := new(mocks.LogoUC)
			handler := &Handler{
				uc:  &usecase.Usecase{Logo: mockLogoUsecase},
				log: logrus.New(),
			}

			tc.mockBehaviour(mockLogoUsecase, tc.args)

			router := gin.Default()
			router.PUT("/logos/:id", handler.updateLogo)

			req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/logos/%v", tc.args.id), bytes.NewBufferString(tc.inputbody))
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatusCode, w.Code)
			assert.JSONEq(t, tc.wantRespBody, w.Body.String())
			mockLogoUsecase.AssertExpectations(t)
		})
	}
}
