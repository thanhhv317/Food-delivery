package common

type Paging struct {
	Limit      int    `json:"limit" form:"limit"`
	Page       int    `json:"page" form:"page"`
	Total      int64  `json:"total" form:"total"`
	FakeCursor string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor"`
}

func (p *Paging) Process() {
	if p.Limit <= 0 {
		p.Limit = 10
	}

	if p.Page <= 0 {
		p.Page = 1
	}
}
