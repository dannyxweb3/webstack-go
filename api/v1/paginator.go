package v1

type Paginator struct {
	CurPage    int `json:"cur_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
	PageSize   int `json:"page_size"`
	PageRange  int `json:"page_range"` // 显示多少个页码按钮
}
