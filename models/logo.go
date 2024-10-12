package models

type Logo struct {
	ID    int    `json:"-"`
	Url   string `json:"url" binding:"required,min=3"`
	Title string `json:"title" binding:"required,min=3,max=100"`
}

type LogoUpdate struct {
	Url   *string `json:"url" binding:"omitempty,min=3"`
	Title *string `json:"title" binding:"omitempty,min=3,max=100"`
}

func (l *LogoUpdate) IsEmpty() bool {
	return l.Url == nil && l.Title == nil
}

type LogoResp struct {
	ID    int    `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
}

type LogosResp struct {
	Logos map[int]*Logo `json:"logos"`
	Total int     `json:"total"`
}
