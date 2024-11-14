package entity

import (
	"net/url"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/lib/convert"
)

type Adv struct {
	ID          int    `json:"id,omitempty"`
	ColorID     int    `json:"color_id"`
	ImageID     int    `json:"logo_id"`
	TgID        int    `json:"tg_id"`
	Post        string `json:"post"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

type AdvUpdate struct {
	ColorID     Optional[int]    `json:"color_id"`
	ImageID     Optional[int]    `json:"logo_id"`
	TgID        Optional[int]    `json:"tg_id"`
	Post        Optional[string] `json:"post"`
	Title       Optional[string] `json:"title"`
	Description Optional[string] `json:"description"`
	Priority    Optional[int]    `json:"priority"`
}

type AdvResp struct {
	ID          int    `json:"id,omitempty"`
	Color       Color  `json:"color"`
	Image       Image  `json:"image"`
	Tg          Tg     `json:"tg"`
	Post        string `json:"post"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

func AdvRespToApi(a *AdvResp) *api.Ad {
	imgUrl, _ := url.Parse(a.Image.Url)
	tgUrl, _ := url.Parse(a.Tg.Url)

	return &api.Ad{
		ID: a.ID,
		Image: api.Image{
			ID:    a.Image.ID,
			Title: a.Image.Title,
			URL:   *imgUrl,
			Type:  api.ImageType(a.Image.Type),
		},
		Color: api.Color{
			ID:    a.Color.ID,
			Title: a.Color.Title,
			Hex:   a.Color.Hex,
		},
		Tg: api.Tg{
			ID:    a.Tg.ID,
			Title: a.Tg.Title,
			URL:   *tgUrl,
		},
		Post:        a.Post,
		Title:       a.Title,
		Description: a.Description,
		Priority:    api.AdPriority(a.Priority),
	}
}

func (a *AdvUpdate) IsSet() bool {
	return a.ColorID.Set || a.ImageID.Set || a.TgID.Set || a.Post.Set || a.Title.Set || a.Description.Set || a.Priority.Set
}

func AdsToApi(ads []*AdvResp) *api.Ads {
	apiAds := api.Ads{}
	for _, ad := range ads {
		apiAds = append(apiAds, *AdvRespToApi(ad))
	}
	
	return &apiAds
}

type AdsResp struct {
	Data    []*AdvResp `json:"data"`
	Page    uint64     `json:"page"`
	Pages   uint64     `json:"pages"`
	PerPage uint64     `json:"per_page"`
}

func (c *AdsResp) ToApi() *api.AdsResp {
	return &api.AdsResp{
		Data:    convert.ApplyPointerToSlice(c.Data, AdvRespToApi),
		Page:    int(c.Page),
		Pages:   int(c.Pages),
		PerPage: int(c.PerPage),
	}
}
