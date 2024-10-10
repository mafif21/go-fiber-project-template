package dtos

type WebResponse[T any] struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    T             `json:"data"`
	Paging  *PageMetadata `json:"paging,omitempty"`
	Errors  string        `json:"errors,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}
