package entity

import (
	"net/url"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/lib/convert"
)

type TgUpdate struct {
	Url   Optional[string] `json:"url" binding:"omitempty,min=3"`
	Title Optional[string] `json:"title" binding:"omitempty,min=3,max=100"`
}

func (t *TgUpdate) IsSet() bool {
	return t.Url.Set || t.Title.Set
}

type Tg struct {
	ID    int    `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
}

func TgToApi(t *Tg) *api.Tg {
	// TODO: ловить ошибку
	url, _ := url.Parse(t.Url)
	return &api.Tg{
		ID:    t.ID,
		Title: t.Title,
		URL:   *url,
	}
}

type TgsResp struct {
	Data    []*Tg  `json:"data"`
	Page    uint64 `json:"page"`
	Pages   uint64 `json:"pages"`
	PerPage uint64 `json:"per_page"`
}

func (c *TgsResp) ToApi() *api.TgsResp {
	return &api.TgsResp{
		Data:    convert.ApplyPointerToSlice(c.Data, TgToApi),
		Page:    int(c.Page),
		Pages:   int(c.Pages),
		PerPage: int(c.PerPage),
	}
}
