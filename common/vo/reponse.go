package vo

type PageResponse struct {
	Total int `json:"total"`
	Size int `json:"size"`
	Current int `json:"current"`
}
