package dto

import (
	"time"
)

type GetTodos struct {
	ID       int       `json:"Id"`
	Title    string    `json:"Title"`
	Content  string    `json:"Content"`
	IsDone   bool      `json:"IsDone"`
	CreateAt time.Time `json:"CreateAt"`
}
