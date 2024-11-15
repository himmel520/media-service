package entity

import (
	"net/url"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/lib/convert"
)

type ImageUpdate struct {
	Url   Optional[string] `json:"url"`
	Title Optional[string] `json:"title"`
	Type  Optional[string] `json:"type"`
}

func (i *ImageUpdate) IsSet() bool {
	return i.Url.Set || i.Title.Set || i.Type.Set
}

type Image struct {
	ID    int    `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

func ImageToApi(i *Image) *api.Image {
	url, _ := url.Parse(i.Url)
	return &api.Image{
		ID:    i.ID,
		Title: i.Title,
		URL:   *url,
		Type:  api.ImageType(i.Type),
	}
}

func ImageToLogoItemApi(i *Image) *api.LogosRespItem {
	url, _ := url.Parse(i.Url)
	return &api.LogosRespItem{
		Title: i.Title,
		URL:   *url,
		Type:  api.LogosRespItemType(i.Type),
	}
}

type LogosResp map[string]*Image

func LogosRespToApi(logos LogosResp) *api.LogosResp {
	apiLogos := api.LogosResp{}

	for id, logo := range logos {
		url, _ := url.Parse(logo.Url)

		apiLogos[id] = api.LogosRespItem{
			Title: logo.Title,
			URL:   *url,
			// TODO: добавить проверку на тип
			Type:  api.LogosRespItemType(logo.Type),
		}
	}

	return &apiLogos
}

type ImagesResp struct {
	Data    []*Image `json:"data"`
	Page    uint64   `json:"page"`
	Pages   uint64   `json:"pages"`
	PerPage uint64   `json:"per_page"`
}

func (c *ImagesResp) ToApi() *api.ImagesResp {
	return &api.ImagesResp{
		Data:    convert.ApplyPointerToSlice(c.Data, ImageToApi),
		Page:    int(c.Page),
		Pages:   int(c.Pages),
		PerPage: int(c.PerPage),
	}
}
