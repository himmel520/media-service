package entity

import (
	"net/url"

	api "github.com/himmel520/media-service/api/oas"
)

type Logo Image

type LogosResp map[string]*Logo

func LogosRespToApi(logos LogosResp) *api.LogosResp {
	apiLogos := api.LogosResp{}

	for id, logo := range logos {
		url, _ := url.Parse(logo.Url)

		apiLogos[id] = api.LogosRespItem{
			Title: logo.Title,
			URL:   *url,
			// TODO: добавить проверку на тип
			Type: api.LogosRespItemType(logo.Type),
		}
	}

	return &apiLogos
}
