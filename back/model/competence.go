package model

type Competence struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Rating int    `json:"rating"`
	Desc   string `json:"desc"`
}
