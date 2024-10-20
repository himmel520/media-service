package entity

type Color struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title" binding:"required,min=3,max=100"`
	Hex   string `json:"hex" binding:"required,min=7,max=7"`
}

type ColorUpdate struct {
	Title *string `json:"title" binding:"omitempty,min=3,max=100"`
	Hex   *string `json:"hex" binding:"omitempty,min=7,max=7"`
}

type ColorResp struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Hex   string `json:"hex"`
}

func (c *ColorUpdate) IsEmpty() bool {
	return c.Hex == nil && c.Title == nil
}

type ColorsResp struct {
	Colors []*ColorResp `json:"colors"`
	Total  int          `json:"total"`
}
