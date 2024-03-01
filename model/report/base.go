package report

type Metrics struct {
	Consume float64 `json:"consume"`
	Ctr     float64 `json:"ctr"`
	Date    string  `json:"date"`
	Hour    int     `json:"hour"`
	Pv      int     `json:"pv"`
}

type SumMetrics struct {
	Consume float64 `json:"consume"`
	Ctr     float64 `json:"ctr"`
	Date    string  `json:"date"`
	Hour    string  `json:"hour"`
	Pv      int     `json:"pv"`
}

type PageInfo struct {
	Page     int `json:"page"`
	Pages    int `json:"pages"`
	PageSize int `json:"page_size"`
	TotalNum int `json:"total_num"`
}
