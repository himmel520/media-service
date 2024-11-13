package entity

import (
	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/lib/convert"
)

type Color struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Hex   string `json:"hex"`
}

func ColorToApi(c *Color) *api.Color {
	return &api.Color{
		ID:    c.ID,
		Title: c.Title,
		Hex:   c.Hex,
	}
}

type ColorUpdate struct {
	Title Optional[string] `json:"title" binding:"omitempty,min=3,max=100"`
	Hex   Optional[string] `json:"hex" binding:"omitempty,min=7,max=7"`
}

func (c *ColorUpdate) IsSet() bool {
	return c.Hex.Set || c.Title.Set
}

type ColorsResp struct {
	Data    []*Color `json:"data"`
	Page    uint64   `json:"page"`
	Pages   uint64   `json:"pages"`
	PerPage uint64   `json:"per_page"`
}

func (c *ColorsResp) ToApi() *api.ColorsResp {
	return &api.ColorsResp{
		Data:    convert.ApplyToSlice(c.Data, ColorToApi),
		Page:    int(c.Page),
		Pages:   int(c.Pages),
		PerPage: int(c.PerPage),
	}
}
