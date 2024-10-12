package dtos

type SearchConf struct {
	Page int `json:"page" validate:"min=1"`
	Size int `json:"size" validate:"min=1,max=100"`
}
