package model

type Article struct {
	ID          int    `json:"id"`
	Titre       string `json:"titre"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Template    string `json:"template"`
}
