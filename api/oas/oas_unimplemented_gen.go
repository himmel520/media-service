// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// V1AdminAdsGet implements GET /v1/admin/ads operation.
//
// Получает список всех реклам с пагинацией.
//
// GET /v1/admin/ads
func (UnimplementedHandler) V1AdminAdsGet(ctx context.Context, params V1AdminAdsGetParams) (r V1AdminAdsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminAdsIDDelete implements DELETE /v1/admin/ads/{id} operation.
//
// Удаляет рекламу по id.
//
// DELETE /v1/admin/ads/{id}
func (UnimplementedHandler) V1AdminAdsIDDelete(ctx context.Context, params V1AdminAdsIDDeleteParams) (r V1AdminAdsIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminAdsIDPut implements PUT /v1/admin/ads/{id} operation.
//
// Обновляет рекламу по id.
//
// PUT /v1/admin/ads/{id}
func (UnimplementedHandler) V1AdminAdsIDPut(ctx context.Context, req *AdPut, params V1AdminAdsIDPutParams) (r V1AdminAdsIDPutRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminAdsPost implements POST /v1/admin/ads operation.
//
// Создает новую рекламу.
//
// POST /v1/admin/ads
func (UnimplementedHandler) V1AdminAdsPost(ctx context.Context, req *AdPost) (r V1AdminAdsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminColorsGet implements GET /v1/admin/colors operation.
//
// Возвращает список цветов с возможностью пагинации.
//
// GET /v1/admin/colors
func (UnimplementedHandler) V1AdminColorsGet(ctx context.Context, params V1AdminColorsGetParams) (r V1AdminColorsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminColorsIDDelete implements DELETE /v1/admin/colors/{id} operation.
//
// Удаляет цвет с указанным id.
//
// DELETE /v1/admin/colors/{id}
func (UnimplementedHandler) V1AdminColorsIDDelete(ctx context.Context, params V1AdminColorsIDDeleteParams) (r V1AdminColorsIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminColorsIDPut implements PUT /v1/admin/colors/{id} operation.
//
// Обновляет цвет с указанным id.
//
// PUT /v1/admin/colors/{id}
func (UnimplementedHandler) V1AdminColorsIDPut(ctx context.Context, req *ColorPut, params V1AdminColorsIDPutParams) (r V1AdminColorsIDPutRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminColorsPost implements POST /v1/admin/colors operation.
//
// Создает новый цвет.
//
// POST /v1/admin/colors
func (UnimplementedHandler) V1AdminColorsPost(ctx context.Context, req *ColorPost) (r V1AdminColorsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminImagesGet implements GET /v1/admin/images operation.
//
// Возвращает список изображений с поддержкой пагинации.
//
// GET /v1/admin/images
func (UnimplementedHandler) V1AdminImagesGet(ctx context.Context, params V1AdminImagesGetParams) (r V1AdminImagesGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminImagesIDDelete implements DELETE /v1/admin/images/{id} operation.
//
// Удаляет изображение с указанным ID.
//
// DELETE /v1/admin/images/{id}
func (UnimplementedHandler) V1AdminImagesIDDelete(ctx context.Context, params V1AdminImagesIDDeleteParams) (r V1AdminImagesIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminImagesIDPut implements PUT /v1/admin/images/{id} operation.
//
// Обновляет данные изображения с указанным ID.
//
// PUT /v1/admin/images/{id}
func (UnimplementedHandler) V1AdminImagesIDPut(ctx context.Context, req *ImagePut, params V1AdminImagesIDPutParams) (r V1AdminImagesIDPutRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminImagesPost implements POST /v1/admin/images operation.
//
// Создает новый элемент изображения.
//
// POST /v1/admin/images
func (UnimplementedHandler) V1AdminImagesPost(ctx context.Context, req *ImagePost) (r V1AdminImagesPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminTgsGet implements GET /v1/admin/tgs operation.
//
// Возвращает список  тг с возможностью пагинации.
//
// GET /v1/admin/tgs
func (UnimplementedHandler) V1AdminTgsGet(ctx context.Context, params V1AdminTgsGetParams) (r V1AdminTgsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminTgsIDDelete implements DELETE /v1/admin/tgs/{id} operation.
//
// Удаляет тг с указанным id.
//
// DELETE /v1/admin/tgs/{id}
func (UnimplementedHandler) V1AdminTgsIDDelete(ctx context.Context, params V1AdminTgsIDDeleteParams) (r V1AdminTgsIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminTgsIDPut implements PUT /v1/admin/tgs/{id} operation.
//
// Обновляет тг с указанным id.
//
// PUT /v1/admin/tgs/{id}
func (UnimplementedHandler) V1AdminTgsIDPut(ctx context.Context, req *TgPut, params V1AdminTgsIDPutParams) (r V1AdminTgsIDPutRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminTgsPost implements POST /v1/admin/tgs operation.
//
// Создает новый тг.
//
// POST /v1/admin/tgs
func (UnimplementedHandler) V1AdminTgsPost(ctx context.Context, req *TgPost) (r V1AdminTgsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdsGet implements GET /v1/ads operation.
//
// Получает список всех реклам с фильтрацией по
// приоритету и должности.
//
// GET /v1/ads
func (UnimplementedHandler) V1AdsGet(ctx context.Context, params V1AdsGetParams) (r V1AdsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1LogosGet implements GET /v1/logos operation.
//
// Возвращает список всех лого.
//
// GET /v1/logos
func (UnimplementedHandler) V1LogosGet(ctx context.Context) (r V1LogosGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}