package model

type Url struct {
	ID       int    `json:"id"`
	Url      string `json:"url"`
	Alias    string `json:"alias"`
	VisitNum int    `json:"visitnum"`
}
