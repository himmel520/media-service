package entity

type Adv struct {
	ID          int    `json:"id,omitempty"`
	ColorID     int    `json:"color_id" binding:"required,min=1"`
	LogoID      int    `json:"logo_id" binding:"required,min=1"`
	TgID        int    `json:"tg_id" binding:"required,min=1"`
	Post        string `json:"post" binding:"required,min=3,max=100"`
	Title       string `json:"title" binding:"required,min=3,max=40"`
	Description string `json:"description" binding:"required,min=10,max=150"`
	Priority    int    `json:"priority" binding:"required,min=1,max=3"`
}

type AdvUpdate struct {
	ColorID     *int    `json:"color_id" binding:"omitempty,min=1"`
	LogoID      *int    `json:"logo_id" binding:"omitempty,min=1"`
	TgID        *int    `json:"tg_id" binding:"omitempty,min=1"`
	Post        *string `json:"post" binding:"omitempty,min=3,max=100"`
	Title       *string `json:"title" binding:"omitempty,min=3,max=40"`
	Description *string `json:"description" binding:"omitempty,min=10,max=150"`
	Priority    *int    `json:"priority" binding:"omitempty,min=1,max=3"`
}

type AdvResponse struct {
	ID          int    `json:"id,omitempty"`
	Color       string `json:"hex"`
	LogoUrl     string `json:"logo_url"`
	TgUrl       string `json:"tg_url"`
	Post        string `json:"post"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

func (a *AdvUpdate) IsEmpty() bool {
	return a.ColorID == nil && a.LogoID == nil && a.TgID == nil && a.Post == nil && a.Title == nil && a.Description == nil && a.Priority == nil
}
