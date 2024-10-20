package httpctrl

type PaginationQuery struct {
	Limit  int `form:"limit,default=20" binding:"omitempty,min=1"`
	Offset int `form:"offset,default=0" binding:"omitempty,min=0"`
}

type AdvPostQuery struct {
	PaginationQuery
	Post     []string `form:"post" binding:"required,dive,min=3,max=100"`
	Priority []string `form:"priority" binding:"omitempty,dive,oneof=1 2 3"`
}

func (a *AdvPostQuery) SetDefaultPriority() {
	a.Priority = []string{"1", "2", "3"}
}
