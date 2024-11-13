package ogen

import (
	"context"

	api "github.com/himmel520/media-service/api/oas"
)

type (
	Handler struct {
		Auth
		Error
		Ad
		Color
		Image
		Tg
	}

	Auth interface {
		HandleAdminBearerAuth(ctx context.Context, operationName string, t api.AdminBearerAuth) (context.Context, error)
	}

	Error interface {
		NewError(ctx context.Context, err error) *api.ErrorStatusCode
	}

	Ad interface {
		V1AdminAdsGet(ctx context.Context, params api.V1AdminAdsGetParams) (api.V1AdminAdsGetRes, error)
		V1AdminAdsIDDelete(ctx context.Context, params api.V1AdminAdsIDDeleteParams) (api.V1AdminAdsIDDeleteRes, error)
		V1AdminAdsIDPut(ctx context.Context, req *api.AdPut, params api.V1AdminAdsIDPutParams) (api.V1AdminAdsIDPutRes, error)
		V1AdminAdsPost(ctx context.Context, req *api.AdPost) (api.V1AdminAdsPostRes, error)
		V1AdsGet(ctx context.Context, params api.V1AdsGetParams) (api.V1AdsGetRes, error)
	}

	Color interface {
		V1AdminColorsGet(ctx context.Context, params api.V1AdminColorsGetParams) (api.V1AdminColorsGetRes, error)
		V1AdminColorsIDDelete(ctx context.Context, params api.V1AdminColorsIDDeleteParams) (api.V1AdminColorsIDDeleteRes, error)
		V1AdminColorsIDPut(ctx context.Context, req *api.ColorPut, params api.V1AdminColorsIDPutParams) (api.V1AdminColorsIDPutRes, error)
		V1AdminColorsPost(ctx context.Context, req *api.ColorPost) (api.V1AdminColorsPostRes, error)
	}

	Image interface {
		V1LogosGet(ctx context.Context) (api.V1LogosGetRes, error)
		V1AdminImagesGet(ctx context.Context, params api.V1AdminImagesGetParams) (api.V1AdminImagesGetRes, error)
		V1AdminImagesIDDelete(ctx context.Context, params api.V1AdminImagesIDDeleteParams) (api.V1AdminImagesIDDeleteRes, error)
		V1AdminImagesIDPut(ctx context.Context, req *api.ImagePut, params api.V1AdminImagesIDPutParams) (api.V1AdminImagesIDPutRes, error)
		V1AdminImagesPost(ctx context.Context, req *api.ImagePost) (api.V1AdminImagesPostRes, error)
	}

	Tg interface {
		V1AdminTgsGet(ctx context.Context, params api.V1AdminTgsGetParams) (api.V1AdminTgsGetRes, error)
		V1AdminTgsIDDelete(ctx context.Context, params api.V1AdminTgsIDDeleteParams) (api.V1AdminTgsIDDeleteRes, error)
		V1AdminTgsIDPut(ctx context.Context, req *api.TgPut, params api.V1AdminTgsIDPutParams) (api.V1AdminTgsIDPutRes, error)
		V1AdminTgsPost(ctx context.Context, req *api.TgPost) (api.V1AdminTgsPostRes, error)
	}
)

type HandlerParams struct {
	Auth
	Error
	Ad
	Color
	Image
	Tg
}

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		Auth:  params.Auth,
		Error: params.Error,
		Ad:    params.Ad,
		Color: params.Color,
		Image: params.Image,
		Tg:    params.Tg,
	}
}
