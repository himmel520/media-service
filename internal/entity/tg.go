package entity

type TG struct {
	ID    int    `json:"id,omitempty"`
	Url   string `json:"url" binding:"required,min=3"`
	Title string `json:"title" binding:"required,min=3,max=100"`
}

type TGUpdate struct {
	Url   *string `json:"url" binding:"omitempty,min=3"`
	Title *string `json:"title" binding:"omitempty,min=3,max=100"`
}

func (t *TGUpdate) IsEmpty() bool {
	return t.Url == nil && t.Title == nil
}

type TGResp struct {
	ID    int    `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
}

type TGsResp struct {
	TGs []*TGResp `json:"tgs"`
	Total  int       `json:"total"`
}
