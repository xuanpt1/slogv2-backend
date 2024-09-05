package temp

type Page struct {
	Page     int   `json:"page" label:"当前页码"`
	PageSize int   `json:"page_size" label:"每页数量"`
	Total    int64 `json:"total" label:"总数量"`
}
