package dto

type UpdateTodo struct {
	Title   string `json:"Title"`
	Content string `json:"Content"`
	IsDone  bool   `json:"IsDone"`
}
