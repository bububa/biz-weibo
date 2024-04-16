package model

type PageInfo struct {
	// Total
	Total int `json:"total,omitempty"`
	// PageTotal
	PageTotal int `json:"page_total,omitempty"`
	// PageSize
	PageSize int `json:"page_size,omitempty"`
	// Page
	Page int `json:"page,omitempty"`
}
