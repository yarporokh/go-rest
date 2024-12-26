package models

type Item struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	UserId uint   `json:"userid"`
}
