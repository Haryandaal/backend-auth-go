package models

type History struct {
	Id     string `json:"id"`
	Action string `json:"action"`
	Detail string `json:"detail"`
	Date   string `json:"date"`
}
